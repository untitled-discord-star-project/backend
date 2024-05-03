package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/scylladb/gocqlx/v2"
	"github.com/untitled-discord-star-project/backend/models"
	"github.com/untitled-discord-star-project/backend/templates"
)

func CreateUiEndpoints(db gocqlx.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")
		switch {
		case strings.HasPrefix(r.URL.Path, "/ui/quote"):
			switch r.Method {
			case http.MethodGet:
				quoteTemplate := createQuoteTemplate(db)

				err := quoteTemplate.Render(r.Context(), w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			default:
				http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			}
		}
	}
}

func createQuoteTemplate(db gocqlx.Session) templ.Component {

	var starboardMessages []models.DiscordStarboardStruct

	q := db.Query(models.DiscordStarboard.SelectAll())

	if err := q.SelectRelease(&starboardMessages); err != nil {
		log.Fatal(err)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rng.Intn(len(starboardMessages))

	return templates.SpeechBubble(starboardMessages[index].MessageContent)

}
