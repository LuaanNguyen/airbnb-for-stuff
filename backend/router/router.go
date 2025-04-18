package router

import (
	"github.com/LuaanNguyen/backend/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/healthcheck", middleware.HealthCheck).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users", middleware.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/items", middleware.GetAllItems).Methods("GET", "OPTIONS")
	return router
}