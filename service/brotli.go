package service

import (
	"net/http"

	"github.com/andybalholm/brotli"
)

func BrotliMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		br := brotli.HTTPCompressor(w, r)
		defer br.Close()
		w.Header().Set("Content-Encoding", "br")
		next.ServeHTTP(w, r)
	})
}
