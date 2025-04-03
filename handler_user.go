package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tshjustin/RSS-aggragator-go/internal/database"
)

// HTTP Handler that handles the creation of a user
// This defines a METHOD (rather than a func) - Such that we can operate on the object apiConfig
func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	// 1. Create a local struct called parameters
	type parameters struct {
		Name string
	}

	// Create a Decoder that reads from the request body
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	// Decodes the actual reqeust body
	// Decode method requires a pointer as an argument. Hence we pass in & to get the mmory adress
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Access the struct
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	// rather than responding with DB user, we respond with our own user
	// respondWithJson(w, 200, user)
	respondWithJson(w, 200, databaseUserToUser(user))
}
