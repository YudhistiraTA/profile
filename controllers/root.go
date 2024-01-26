package controllers

import (
	"net/http"

	"github.com/YudhistiraTA/profile/lib"
	"github.com/YudhistiraTA/profile/views/components"
	"github.com/YudhistiraTA/profile/views/layouts"
)

func Root(w http.ResponseWriter, r *http.Request) {
	lib.Htmx(w, r, layouts.Main(components.Main()))
}
