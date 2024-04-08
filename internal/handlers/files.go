package handlers

import (
	"net/http"
)

func FilesEndpoint(w http.ResponseWriter, r *http.Request) {
	switch (r.Method) {
	case http.MethodGet:
		fileServer := http.FileServer(http.Dir("."))
		fileServer.ServeHTTP(w, r)
	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
