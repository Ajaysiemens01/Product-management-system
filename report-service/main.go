package main

import (
	"log"
	"net/http"
	"report-service/handlers"
	"context"
	"os/signal"
	"syscall"
	"time"
	"github.com/gorilla/mux"
	"report-service/middlewares"
	"report-service/config"
)



func main() {
    r := mux.NewRouter()

	config.LoadConfig() 
	
    // // Apply authentication middleware
	 api := r.PathPrefix("/api").Subrouter()
	 api.Use(middlewares.APIKeyMiddleware)

    // Define routes
	api.HandleFunc("/report/inventory", handlers.GetInventoryReportHandler).Methods("GET")

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
	// Make a best effort to shut down the server cleanly. We don’t
	// need to collect the server’s error if we didn’t already;
	// Shutdown will let us know (unless something worse happens, in
	// which case it will tell us that).
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
