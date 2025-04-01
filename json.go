package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	// Transalate Go data to Json Data
	dat, err := json.Marshal(payload) // This returns as bytes - to allow bin format
	if err != nil {
		log.Printf("Failed to marshal Json Reponse %v", payload)
		w.WriteHeader(500)
		return
	}

	// like writing on the envelope "This contains JSON" - Tells the reciepient what kind of data to accept , in this case json
	w.Header().Add("Content-Type", "application/json")

	// Stamping the envelope with a status code
	w.WriteHeader(code)

	// FIlling the envelop with a letter (json data)
	w.Write(dat) // writes the

	// NOTE that => w => Envelop (resposnse package) | data = letter (actual payload)
}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errorResponse{
		Error: msg,
	})
}