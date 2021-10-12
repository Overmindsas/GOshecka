package server

import "net/http"

type responsWriter struct {
	http.ResponseWriter
	code int
}

func (w *responsWriter) WriteHeader(status int) {
	w.code = status
	w.ResponseWriter.WriteHeader(status)
}
