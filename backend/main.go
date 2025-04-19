package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LuaanNguyen/backend/db"
	"github.com/LuaanNguyen/backend/router"
)

func main() {
	// Initialize database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.DB.Close()

	// Create router with database connection
	r := router.Router(db.DB)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}