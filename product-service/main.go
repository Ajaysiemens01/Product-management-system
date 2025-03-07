package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"product-service/handler"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"product-service/config"
	"product-service/middlewares"
)

func main() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "X-API-KEY"},
		AllowCredentials: true,
	})

	config.LoadConfig()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.APIKeyMiddleware)

	// Define routes
	api.HandleFunc("/products", handler.GetProductsHandler).Methods("GET")
	api.HandleFunc("/products", handler.AddProductHandler).Methods("POST")

	portString := config.PORT
	if portString == "" {
		log.Fatal("Port Not found in the environment")
	}
	log.Println("Server started on : " + portString)

	server := &http.Server{
		Addr:    ":" + portString,
		Handler: c.Handler(r),
	}

	// Graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	serr := make(chan error, 1)

	go func() { serr <- server.ListenAndServe() }()

	var e error
	select {
	case e = <-serr:
	case <-ctx.Done():
	}

	sdctx, sdcancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer sdcancel()
	if shutdownErr := server.Shutdown(sdctx); shutdownErr != nil {
		log.Printf("Server Shutdown Failed:%+v", shutdownErr)
	}
	if e != nil {
		log.Printf("Server encountered an error: %+v", e)
	} else {
		log.Println("Server shut down gracefully.")
	}
}
