package server

import (
	"net/http"
)

// Because while developing ReactJS we might have CORS issues as the development
// server and our Golang server are going to run in different ports, we
// might need to tell our server which Origins are valid here.
func handleCORS(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	validDomains := []string{
		"http://localhost:8080",
		"http://localhost:3000",
	}

	if origin == validDomains[0] || origin == validDomains[1] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cookie",
		)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	}
}
