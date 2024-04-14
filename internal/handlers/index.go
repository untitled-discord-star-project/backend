package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func CreateIndexEndpoint(t templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Print("IndexEndpoint")
			err := t.Render(r.Context(), w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
