package view

import (
	"net/http"

	"github.com/YudhistiraTA/profile/view/template"
	"github.com/a-h/templ"
)

func Htmx(w http.ResponseWriter, r *http.Request, component templ.Component) {
	templ.Handler(template.Index(component)).ServeHTTP(w, r)
}
