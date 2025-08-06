package user

import (
	"time"

	"github.com/google/uuid"
)

/*
import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5" // go get -u github.com/golang-jwt/jwt/v5
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt" // "go get golang.org/x/crypto/bcrypt" //to install
)

*/

type User struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	IsChirpyRed  bool      `json:"is_chirpy_red"`
}
