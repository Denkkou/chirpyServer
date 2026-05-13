package main

import (
	"net/http"
)

func main() {
	// Create new serveMux
	mux := http.NewServeMux()

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()

}
