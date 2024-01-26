package lib

import (
	"net/http"

	"github.com/YudhistiraTA/profile/views"

	"github.com/a-h/templ"
)

func Htmx(w http.ResponseWriter, r *http.Request, component templ.Component) {
	templ.Handler(views.Index(component)).ServeHTTP(w, r)
}
