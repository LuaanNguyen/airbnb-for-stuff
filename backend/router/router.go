package router

import (
	"database/sql"

	"github.com/LuaanNguyen/backend/handlers"
	"github.com/LuaanNguyen/backend/middleware"
	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.EnableCORS) // Apply CORS middleware globally

	//  -------------- Public routes (no auth required)  --------------
	router.HandleFunc("/healthcheck", handlers.HealthCheck).Methods("GET", "OPTIONS")
	router.HandleFunc("/login", handlers.Login).Methods("POST", "OPTIONS")

	// -------------- Protected routes with /api/ prefix  --------------
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)  // Auth only for protected routes
	
	// User routes
	protected.HandleFunc("/users", handlers.GetAllUser).Methods("GET", "OPTIONS")
	protected.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET", "OPTIONS")

	// Item routes
	protected.HandleFunc("/items", handlers.GetAllItems).Methods("GET", "OPTIONS")
	protected.HandleFunc("/items", handlers.CreateItem).Methods("POST", "OPTIONS")
	protected.HandleFunc("/items/available", handlers.GetAvailableItems).Methods("GET", "OPTIONS")
	protected.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET", "OPTIONS")
	protected.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT", "OPTIONS")
	protected.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE", "OPTIONS")
	protected.HandleFunc("/items/search", handlers.SearchItems).Methods("GET", "OPTIONS")

	// Rental routes
	protected.HandleFunc("/rentals", handlers.CreateRentalRequest).Methods("POST", "OPTIONS")
	protected.HandleFunc("/rentals/my", handlers.GetMyRentals).Methods("GET", "OPTIONS")

	// Category routes
	protected.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET", "OPTIONS")

	// Transaction routes
	// protected.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")
	// protected.HandleFunc("/user/{id}/transactions", handlers.GetUserTransactions)
	// protected.HandleFunc("/transactions/{id}", handlers.GetTransaction)
	// protected.HandleFunc("/transactions/{id}", handlers.UpdateTransaction).Methods("PUT")

	// Review routes
	// protected.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	// protected.HandleFunc("/items/{id}/reviews", handlers.GetItemReviews)
	// protected.HandleFunc("/user/{id}/reviews", handlers.GetUserReviews)
	// protected.HandleFunc("/reviews/{id}", handlers.UpdateReview).Methods("PUT")
	// protected.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")

	return router
}