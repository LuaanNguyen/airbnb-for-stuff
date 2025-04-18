package main

import (
	"log"
	"net/http"

	"github.com/LuaanNguyen/backend/router"
)

func main() {
	router := router.Router()
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}