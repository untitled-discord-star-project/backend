package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/untitled-discord-star-project/backend/internal/handlers"
	"github.com/untitled-discord-star-project/backend/pkg/middleware"
	"github.com/untitled-discord-star-project/backend/templates"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	db := initDatabase()
	defer db.Close()

	indexTemplate := templates.Index("2so cool!!")
	indexEndpoint := handlers.CreateIndexEndpoint(indexTemplate)
	UiEndpoint := handlers.CreateUiEndpoints(*db)
	APIEnpoints := handlers.CreateAPIEndpoints(*db)

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/":
			indexEndpoint(w, r)
		case strings.HasPrefix(r.URL.Path, "/static"):
			handlers.FilesEndpoint(w, r)
		case strings.HasPrefix(r.URL.Path, "/ui"):
			UiEndpoint(w, r)
		case strings.HasPrefix(r.URL.Path, "/api"):
			APIEnpoints(w, r)
		case r.URL.Path == "/message":
			handlers.MessageEndpoint(w, r)
		default:
			http.NotFound(w, r)
		}
	}

	handler = middleware.CompressionMiddleware(handler)
	handler = middleware.RecordResponse(handler)
	handler = middleware.Recovery(handler)
	handler = middleware.PermissiveCORSHandler(handler)
	handler = middleware.Log(handler)
	handler = middleware.Trace(handler)

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", *port),
		Handler:           handler,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		ReadHeaderTimeout: 200 * time.Millisecond,
	}

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		slog.Error(err.Error())
	}
}

func initDatabase() *gocqlx.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "discord"
	cluster.Consistency = gocql.Quorum
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}

	return &session
}
