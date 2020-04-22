package main

import (
	"net/http"
)

// Test - controller for GET /
func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Server is active..."))
}
