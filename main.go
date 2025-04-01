package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/tshjustin/RSS-aggragator-go/internal/database"

	_ "github.com/lib/pq" // Includes the code in the program although not using directly - This is a database driver

)

type apiConfig struct {
	DB * database.Queries
}

func main() {

	// Loads the env first
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("No DB URL Found")
	}

	// Connect to the database now 
	conn ,err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Connect Connect to the Database", err)
	}

	// Convert sql.queries to what is needed by Open() 
	// Now with this, if we pass to some handler, they can access the DB 
	apiCfg := apiConfig {
		DB: database.New(conn),
	}

	// +--------------+
	// | Router setup |
	// +--------------+

	// A router decides where to send web requests based on their URLs
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


	// Sets up a v1 router, that is mounted to the main router -> v1/healthz would invoke the handlerReadiness 
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) 
	v1Router.Get("/err", handlerError) 
	v1Router.Post("/users", apiCfg.handlerCreateUser) // Now the handler has access to the DB - Note how it becomes a METHDO instead, and we access via the dot operator
	router.Mount("/v1", v1Router)


	// +--------------+
	// | Server setup |
	// +--------------+

	// Set up server that does the following:
	// 1. Use our router to handle incoming reqs
	// 2. Address where port would listen to
	srv := &http.Server{ // Pointer here since the server object is large, thus using a pointer to point to it is much more eff - handles HTTP reqs and serve web content (http server)
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on %v", portString)

	// Start the server
	err2 := srv.ListenAndServe() // This would allow actual changes to the server when using a pointer
	if err2 != nil {
		log.Fatal(err2)
	}

}
