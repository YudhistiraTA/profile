package service

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/YudhistiraTA/profile/controller"
	"github.com/YudhistiraTA/profile/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func FileServer(r chi.Router) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "service", "static"))
	fileServer := http.StripPrefix("/static", http.FileServer(filesDir))
	staticRouter := chi.NewRouter()
	staticRouter.Use(BrotliMiddleware)
	staticRouter.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})
	r.Mount("/static", staticRouter)
}

func NewServer(Addr string, d *db.Database, rc *db.RedisClient) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger, CORS)

	// static dir serve
	FileServer(r)

	// routes
	c := controller.NewController(d, rc)
	r.Get("/", c.Root)
	r.Get("/md", c.Md)
	r.Get("/{fileName}", c.MdPage)

	return &http.Server{
		Addr:    Addr,
		Handler: r,
	}
}
