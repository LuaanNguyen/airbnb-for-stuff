package middleware

import "net/http"

func EnableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set specific origin instead of wildcard '*' for better security
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        // Important headers for CORS
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        w.Header().Set("Access-Control-Allow-Credentials", "true") // For cookies if needed
        w.Header().Set("Access-Control-Max-Age", "3600") // Cache preflight for 1 hour
        
        // Handle preflight requests (OPTIONS)
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
} 