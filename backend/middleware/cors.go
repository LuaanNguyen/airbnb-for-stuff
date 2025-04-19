package middleware

import "net/http"

func EnableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any domain (For development only)
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Specifies what HTTP are allowed
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allows headers like Content-Type and Authorization
        
        if r.Method == "OPTIONS" { // handle "preflights" requests
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
} 