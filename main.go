package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Loads the env first
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// A router decides where to send web requests based on their URLs.
	router := chi.NewRouter()

	// Adding CORS configuration to allow reqs from browser
	// This would allow the server to send more http headers in response - Tells browser "hey you can send from https / http | use this methods etc -> Note that this is a SERVER response"
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// We want to mount this to another path
	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)

	// Set up server that does the following:
	// 1. Use our router to handle incoming reqs
	// 2. Address where port would listen to
	srv := &http.Server{ // Pointer here since the server object is large, thus using a pointer to point to it is much more eff - handles HTTP reqs and serve web content (http server)
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on %v", portString)

	// Start the server
	err := srv.ListenAndServe() // This would allow actual changes to the server when using a pointer
	if err != nil {
		log.Fatal(err)
	}

}
