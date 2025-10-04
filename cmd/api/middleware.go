package main

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type wrap struct {
	http.ResponseWriter
	status int
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := newWrap(w)
		next.ServeHTTP(ww, r)
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", ww.status).
			Dur("duration_ms", time.Since(start)).
			Msg("request completed")
	})
}

func newWrap(w http.ResponseWriter) *wrap {
	return &wrap{ResponseWriter: w, status: http.StatusOK}
}

func (w *wrap) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
