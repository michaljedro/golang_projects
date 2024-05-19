package main

import (
	"log"
	"myapp/handlers"
	"myapp/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    // Initialize the database
    db, err := utils.NewDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Initialize the router
    r := mux.NewRouter()

    // Pass the database to handlers
    handlers.InitializeHandlers(r, db)

    // Start the server
    log.Println("Server running on port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
