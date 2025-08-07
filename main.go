package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" //needed for pq to work correctly
)

func main() {
	dbURL := os.Getenv("DB_URL")           // loads environmental variables from the .env in root,
	db, err := sql.Open("postgres", dbURL) // opens connection to the database
	/*
		Use your SQLC generated database package to create a new *database.Queries,
		 and store it in your apiConfig struct so that handlers can access it:

	*/

	dbQueries := database.New(db)

	fmt.Printf("Welcome to myTop3!")

}
