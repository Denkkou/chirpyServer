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

	// Create new apiConfig
	cfg := &apiConfig{}

	// Standard FileServer as handler
	mux.Handle(appPath, http.StripPrefix(appPath, cfg.middlewareMetricsInc(http.FileServer(http.Dir(rootPath)))))

	// Register endpoints
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)

	// Set server port and handler
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start server
	log.Printf("Serving files from %s on port: %s\n", appPath, port)
	log.Fatal(srv.ListenAndServe()) // Main blocks until server shuts down
}
