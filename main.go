package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" //needed for pq to work correctly
)

type apiConfig struct {
	db *database.Queries
}

type errResponse struct {
	Error string `json:"error"`
}

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	dbURL := os.Getenv("DB_URL")           // loads environmental variables from the .env in root,
	db, err := sql.Open("postgres", dbURL) // opens connection to the database
	if err != nil {
		fmt.Println("error opening sql: ", err)
		os.Exit(1)
	}
	defer db.Close()

	/*
		Use your SQLC generated database package to create a new *database.Queries,
		 and store it in your apiConfig struct so that handlers can access it:

	*/
	dbQueries := database.New(db) // waiting on sqlc to generate this

	/*
		platform := os.Getenv("PLATFORM")
		secret := os.Getenv("SECRET")
		polkaKey := os.Getenv("POLKA_KEY")
	*/

	fmt.Printf("Welcome to myTop3!")
	cfg := &apiConfig{
		db: dbQueries,
		//platform: platform,
		//secret:   secret,
		//polkaKey: polkaKey,
	}

	// This creates a "multiplexer"—a router for incoming HTTP requests.
	// It decides which handler should process requests for different URL paths.
	mux := http.NewServeMux()

	// Actually makes the server that listens on port 8080 and uses the mux that was just created.
	newServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// Tells tbe mux that any request starting with "/" should be handled by a fileserver serving from
	// the current directory.
	//  This allows files like "index.html" (and other static files) to be served for most requests.
	// first version:
	// mux.Handle("/", http.FileServer(http.Dir(".")))
	// after adding readiness():
	mux.Handle("POST /api/users", cfg.createUser)

	err = newServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func (cfg *apiConfig) createUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	newUserParams := CreateUserParams{}
	err := decoder.Decode(&newUserParams)
	if err != nil {
		respondWithError(w, 500, "Error decoding new user params")
		return
	}

	newUserParams.Password, err = auth.HashPassword(newUserParams.Password)
	if err != nil {
		respondWithError(w, 500, "Error creating user password")
		return
	}

	//*************************************************************
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	resp := errResponse{Error: msg}
	jsonWriter(w, code, resp)
}

func jsonWriter(w http.ResponseWriter, code int, payload interface{}) {

	jsonBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Printf("error marshalling response: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError) // auto handles setting header to 500 and body to error
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonBytes)
}
