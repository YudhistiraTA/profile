package lib

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func FileServer(r chi.Router) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fileServer := http.StripPrefix("/static", http.FileServer(filesDir))

	r.Get("/static"+"/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})
}
