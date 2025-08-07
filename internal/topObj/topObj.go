package topobj

import (
	"time"

	"github.com/google/uuid" //go get github.com/google/uuid
)

type topObj struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Name       string    `json:"top_obj_name"`
	SearchKeys []string  `json:"search_keys"`

	UserID uuid.UUID `json:"user_id"`
}

/*

games table: id, display_name, search_key

*/
