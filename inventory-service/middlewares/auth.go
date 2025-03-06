package middlewares

import (
	"log"
	"net/http"
	"inventory-service/config"
	"time"
)

// API key authentication middleware with CORS
func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins (change for production)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-KEY")

		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// API key authentication
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" || apiKey != config.APIKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Log request and measure time
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
