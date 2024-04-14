package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/scylladb/gocqlx/v2"
	"github.com/untitled-discord-star-project/backend/models"
)

func CreateAPIEndpoints(db gocqlx.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/starboard"):
			switch r.Method {
			case http.MethodPost:
				var starboardMessage models.DiscordStarboardStruct
				if err := json.NewDecoder(r.Body).Decode(&starboardMessage); err != nil {
					http.Error(w, fmt.Sprintf("400 Bad Request: %v", err), http.StatusBadRequest)
					log.Println(err)
					return
				}

				query := db.Query(models.DiscordStarboard.Insert()).BindStruct(starboardMessage)

				if err := query.ExecRelease(); err != nil {
					http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
					log.Println(err)
					return
				}

				w.WriteHeader(http.StatusCreated)

			default:
				http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			}
		}
	}

}
