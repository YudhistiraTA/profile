package services

import (
	"net/http"

	"github.com/YudhistiraTA/profile/controllers"
	"github.com/YudhistiraTA/profile/db"
	"github.com/YudhistiraTA/profile/lib"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewServer(Addr string, d *db.Database, rc *db.RedisClient) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// static dir serve
	lib.FileServer(r)

	// routes
	c := controllers.NewController(d, rc)
	r.Get("/", c.Root)
	r.Get("/md", c.Md)
	r.Get("/{fileName}", c.MdPage)

	return &http.Server{
		Addr:    Addr,
		Handler: r,
	}
}
