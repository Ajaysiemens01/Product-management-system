package middlewares

import (
	"inventory-service/config"
	"log"
	"net/http"
	"time"
)

// API key authentication middleware
func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" || apiKey != config.APIKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
