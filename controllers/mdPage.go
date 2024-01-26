package controllers

import (
	"net/http"
	"os"

	"github.com/YudhistiraTA/profile/lib"
	"github.com/YudhistiraTA/profile/views/components"
	"github.com/YudhistiraTA/profile/views/layouts"
	"github.com/go-chi/chi"
)

func MdPage(w http.ResponseWriter, r *http.Request) {
	hxRequest := r.Header.Get("HX-Request")
	fileName := chi.URLParam(r, "fileName")
	body, toc, err := lib.MdParse("data/" + fileName + ".md")
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
			return
		}
	}
	render := components.MdPage(fileName, body, toc)
	if hxRequest == "true" {
		lib.Htmx(w, r, render)
	} else {
		lib.Htmx(w, r, layouts.Main(render))
	}
}
