package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/untitled-discord-star-project/backend/internal/handlers"
	"github.com/untitled-discord-star-project/backend/pkg/middleware"
	"github.com/untitled-discord-star-project/backend/templates"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	indexTemplate := templates.Index("2so cool!!")
	indexEndpoint := handlers.CreateIndexEndpoint(indexTemplate)

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/":
			indexEndpoint(w, r)
		case strings.HasPrefix(r.URL.Path, "/static"):
			handlers.FilesEndpoint(w, r)
		case r.URL.Path == "/message":
			handlers.MessageEndpoint(w, r)
		default:
			http.NotFound(w, r)
		}
	}

	handler = middleware.RecordResponse(handler)
	handler = middleware.Recovery(handler)
	handler = middleware.PermissiveCORSHandler(handler)
	handler = middleware.Log(handler)
	handler = middleware.Trace(handler)

	server := http.Server {
		Addr: fmt.Sprintf(":%d", *port),
		Handler: handler,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadHeaderTimeout: 200 * time.Millisecond,
	}

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		slog.Error(err.Error())
	}
}
