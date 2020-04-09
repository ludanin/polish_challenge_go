package server

import (
	"encoding/json"
	"net/http"
	"polish/data"
	"strings"
)

// Run starts the server
func Run() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// We're ignoring all other methods except for GET, especially because
		// AJAX requests tends to send a request with the method OPTIONS to
		// our server
		if r.Method == http.MethodGet {
			handleCORS(w, r)

			// Anything that doesn't contains this in its URI is probably a static
			// file
			if strings.Contains(r.RequestURI, "api/get") {
				people := data.GetPeople()
				response, err := json.Marshal(people)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.Write(response)
				}
			} else if r.RequestURI == "/" {
				http.ServeFile(w, r, "./build/index.html")
			} else {
				http.ServeFile(
					w,
					r,
					("./build/" + strings.TrimPrefix(r.RequestURI, "/")),
				)
			}
		}
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
