package main

import (
	"log"
	"net/http"
	"product-service/handlers"
	"context"
	"os/signal"
	"syscall"
	"time"
	"github.com/gorilla/mux"
	"product-service/middlewares"
	"product-service/config"
)

func main() {
    r := mux.NewRouter()

	config.LoadConfig() 
    // Apply authentication middleware
	 api := r.PathPrefix("/api").Subrouter()
	 api.Use(middlewares.CORSMiddleware)
	api.Use(middlewares.APIKeyMiddleware)

	// Define routes
	api.HandleFunc("/products", handlers.GetProductsHandler).Methods("GET")
	api.HandleFunc("/products", handlers.AddProductHandler).Methods("POST")

    portString := config.PORT
    if portString == "" {
        log.Fatal("Port Not found in the environment")
    }
    log.Println("Server started on : " + portString)

    server := &http.Server{
  
        Addr:    ":" + portString,
        Handler: r,
    }
   // When this context is canceled, we will gracefully stop the server.
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	// When the server is stopped *not by that context*, but by any
	// other problems, it will return its error via this.
	serr := make(chan error, 1)

	// Start the server and collect its error return.
	go func() { serr <- server.ListenAndServe() }()

	// Wait for either the server to fail, or the context to end.
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
