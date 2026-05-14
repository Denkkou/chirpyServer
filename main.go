package main

import (
	"log"
	"net/http"
)

func main() {
	const rootPath = "."
	const appPath = "/app/"
	const port = "8080"

	// Create new serveMux
	mux := http.NewServeMux()

	// Standard FileServer as handler
	mux.Handle(appPath, http.StripPrefix(appPath, http.FileServer(http.Dir(rootPath))))

	// Register readiness endpoint
	mux.HandleFunc("/healthz", readinessEndpoint)

	// Set server port and handler
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start server
	log.Printf("Serving files from %s on port: %s\n", appPath, port)
	log.Fatal(srv.ListenAndServe()) // Main blocks until server shuts down
}

func readinessEndpoint(w http.ResponseWriter, req *http.Request) {
	// Write header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Write status code
	w.WriteHeader(http.StatusOK)

	// Write body
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
