package router

import (
	"database/sql"

	"github.com/LuaanNguyen/backend/handlers"
	"github.com/LuaanNguyen/backend/middleware"
	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Apply CORS middleware globally
	router.Use(middleware.EnableCORS)

	// Public routes (no auth required)
	router.HandleFunc("/healthcheck", handlers.HealthCheck)
	router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/users", handlers.GetAllUser)

	// Protected routes with /api/ prefix
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)  // Auth only for protected routes
	
	// User routes
	// protected.HandleFunc("/users", handlers.GetAllUser)
	// protected.HandleFunc("/user/{id}", handlers.GetUser)
	// protected.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")

	// Item routes
	// protected.HandleFunc("/items", handlers.GetAllItems)
	// protected.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	// protected.HandleFunc("/items/{id}", handlers.GetItem)
	// protected.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	// protected.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")
	// protected.HandleFunc("/items/search", handlers.SearchItems)
	// protected.HandleFunc("/items/available", handlers.GetAvailableItems)

	// Category routes
	//protected.HandleFunc("/categories", handlers.GetAllCategories)
	//protected.HandleFunc("/categories/{id}/items", handlers.GetItemsByCategory)

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