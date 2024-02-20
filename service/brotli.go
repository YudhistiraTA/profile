package service

import (
	"net/http"

	"github.com/andybalholm/brotli"
)

func BrotliMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		br := brotli.HTTPCompressor(w, r)
		br.Close()
		next.ServeHTTP(w, r)
	})
}
