package main

import "net/http"

// HTTP Handler to return json

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 400, "Something went wrong")
}
