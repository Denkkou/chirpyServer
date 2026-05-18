package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	// Write header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Write status code
	w.WriteHeader(http.StatusOK)

	// Write body
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
