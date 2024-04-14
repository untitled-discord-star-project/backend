package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/scylladb/gocqlx/v2"
	"github.com/untitled-discord-star-project/backend/models"
	"github.com/untitled-discord-star-project/backend/templates"
)

func StarboardEndpoint(w http.ResponseWriter, r *http.Request, db *gocqlx.Session) {
	switch r.Method {
	case http.MethodGet:
		var starboardMessages []models.DiscordStarboardStruct

		q := db.Query(models.DiscordStarboard.SelectAll())

		if err := q.SelectRelease(&starboardMessages); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		if r.Header.Get("HX-Request") == "true" {
			w.Header().Set("Content-Type", "text/html")

			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			index := rng.Intn(len(starboardMessages))

			randomMessage := starboardMessages[index].MessageContent
			speechBubbleComponent := templates.SpeechBubble(randomMessage)
			err := speechBubbleComponent.Render(r.Context(), w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&starboardMessages); err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

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
