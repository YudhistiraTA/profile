package main

import (
	"net/http"

	"github.com/YudhistiraTA/profile/controllers"
	"github.com/YudhistiraTA/profile/lib"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// chi init
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// static dir serve
	lib.FileServer(r)

	// routes
	r.Get("/", controllers.Root)
	r.Get("/md", controllers.Md)
	r.Get("/{fileName}", controllers.MdPage)

	// start server
	http.ListenAndServe(":3000", r)
}
