package router

import (
	"github.com/LuaanNguyen/backend/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", middleware.HealthCheck).Methods("GET", "OPTIONS")

	router.HandleFunc("/users", middleware.GetAllUser).Methods("GET", "OPTIONS")

	return router
}