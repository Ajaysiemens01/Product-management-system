package middlewares

import (
	"net/http"
	"time"
	"log"
	"product-service/config"
)

func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins (for dev; restrict in production)
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-KEY, Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true") // Required for credentials
        // Handle preflight (OPTIONS) request
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent) // 204 No Content is better for preflight
            return
        }

        next.ServeHTTP(w, r)
    })
}

// API key authentication middleware
func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" || apiKey != config.APIKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		start := time.Now()
		next.ServeHTTP(w,r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}