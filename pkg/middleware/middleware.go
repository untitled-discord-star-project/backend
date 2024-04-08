package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/untitled-discord-star-project/backend/pkg/ctxutil"
	"github.com/untitled-discord-star-project/backend/pkg/trace"

	"github.com/google/uuid"
)

func Trace(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context();

		traceID, err := uuid.Parse(r.Header.Get("X-Trace-ID"))
		if err != nil {
			traceID = uuid.New()
		}

		reqID, err := uuid.Parse(r.Header.Get("X-Request-ID"))
		if err != nil {
			reqID = uuid.New()
		}

		trace := trace.Trace{TraceID: traceID, RequestID: reqID}
		ctx = ctxutil.WithValue(ctx, trace)
		r = r.WithContext(ctx)

		r.Header.Set("X-Trace-ID", trace.TraceID.String())
		r.Header.Set("X-Request-ID", trace.RequestID.String())

		next.ServeHTTP(w, r) 
	}
}

func Log(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var slogger = slog.New(slog.NewTextHandler(os.Stdout, nil)).With(
			slog.String("Method", r.Method),
			slog.Any("URL", r.URL),
		)

		trace, ok := ctxutil.Value[trace.Trace](r.Context())
		if ok {
			slogger = slogger.With(
				slog.Any("TraceID", trace.TraceID),
				slog.Any("RequestID", trace.RequestID),
			)
		}

		ctx := ctxutil.WithValue(r.Context(), slogger)
		next.ServeHTTP(w, r.Clone(ctx))
	}
}

func PermissiveCORSHandler(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func RecordResponse(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		rrw := &RecordingResponseWriter{ ResponseWriter: w }
		start := time.Now()

		next.ServeHTTP(rrw, r)
		elapsed := time.Since(start)

		slogger, ok := ctxutil.Value[*slog.Logger](r.Context())

		if ok {
			slogger.Info("LOG", "StatusCode", rrw.StatusCode, "Status", http.StatusText(rrw.StatusCode), "Bytes", rrw.Bytes, "Elapsed", elapsed)
		} else {
			slog.Info("DEFAULT", "StatusCode", rrw.StatusCode, "Status", http.StatusText(rrw.StatusCode), "Bytes", rrw.Bytes, "Elapsed", elapsed)
		}
	}
}

func Recovery(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				slogger, ok := ctxutil.Value[*slog.Logger](r.Context())
				if ok {
					slogger.Error("LOG", "Panic", err, "Stack", stack)
				} else {
					slog.Error("DEFAULT", "Panic", err, "Stack", stack)
				}

				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
			}
		}()
		next.ServeHTTP(w, r)
	}
}
