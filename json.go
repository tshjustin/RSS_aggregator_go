package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	// interface -> allows us to morph to json
	dat, err := json.Marshal(payload) // This returns as bytes - to allow bin format
	if err != nil {
		log.Printf("Failed to marshal Json Reponse %v", payload)
		w.WriteHeader(500)
		return
	}

	// Adds response headers with content type of return = json
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat) // writes the data

}
