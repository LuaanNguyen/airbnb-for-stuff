package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LuaanNguyen/backend/middleware"
	"github.com/gorilla/mux"
)

func main() {
	// create a new router 
	router := mux.NewRouter()


	// Healthcheck route 
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		db := middleware.CreateConnection() // establish connection with 
		defer db.Close()

		resp := middleware.Response {
			Message: "Hello, you have successfully connected to Postgres 🫶",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Get all users route
	router.HandleFunc("/users", middleware.GetAllUser).Methods("GET", "OPTIONS")

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}