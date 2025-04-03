// Instead of returning the shcema of the db object, we add our own type / config and then return that 

package main

import (
	"time"

	"github.com/tshjustin/RSS-aggragator-go/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"` // <fieldname> <type> <json Tag> (When converted to JSON, this field will be named "id")
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

// Take a database user -> Convert to the schema above 
func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
	}
}
