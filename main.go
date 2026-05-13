package main

import (
	"log"
	"net/http"
)

const root = "/"
const port = "8080"

func main() {
	// Create new serveMux
	mux := http.NewServeMux()

	// Standard FileServer as handler for root path
	mux.Handle(root, http.FileServer(http.Dir(".")))

	// Set server port and handler
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start server
	log.Printf("Serving files from %s on port: %s\n", root, port)
	log.Fatal(srv.ListenAndServe()) // Main blocks until server shuts down
}
