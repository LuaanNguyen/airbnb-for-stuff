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
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		db := middleware.CreateConnections()
		defer db.Close()

		resp := middleware.Response {
			Message: "Hello, you have successfully connected to Postgres 🫶",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}