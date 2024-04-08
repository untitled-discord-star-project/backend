package middleware

import "net/http"

type RecordingResponseWriter struct {
	ResponseWriter http.ResponseWriter
	StatusCode int
	Bytes int
}

func (w *RecordingResponseWriter) WriteHeader(statusCode int) {
	if (w.StatusCode == 0) {
		w.StatusCode = statusCode
	}

	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *RecordingResponseWriter) Header() http.Header { return w.ResponseWriter.Header() }

func (w *RecordingResponseWriter) Write(b []byte) (int, error) {
	if w.StatusCode == 0 {
		w.WriteHeader(http.StatusOK)
	}

	n, err := w.ResponseWriter.Write(b)
	w.Bytes += n
	
	return n, err
}
