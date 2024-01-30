package controllers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/YudhistiraTA/profile/db"
	"github.com/YudhistiraTA/profile/lib"
	"github.com/YudhistiraTA/profile/views/components"
	"github.com/YudhistiraTA/profile/views/layouts"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Controller struct {
	DB *db.Database
}

func NewController(DB *db.Database) *Controller {
	return &Controller{
		DB: DB,
	}
}

const (
	internalServerErrorMsg = "500 - Internal Server Error"
	notFoundErrorMsg       = "404 - Not Found"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, msg string, err error) {
	log.Println(err)
	http.Error(w, msg, statusCode)
}

func (c *Controller) Root(w http.ResponseWriter, r *http.Request) {
	lib.Htmx(w, r, layouts.Main(components.Main()))
}
func (c *Controller) Md(w http.ResponseWriter, r *http.Request) {
	dir := "./data"
	files, err := os.ReadDir(dir)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, internalServerErrorMsg, err)
		return
	}
	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filename := filepath.Base(file.Name())
			filename = strings.TrimSuffix(filename, filepath.Ext(filename))
			filenames = append(filenames, filename)
		}
	}
	w.Write([]byte(strings.Join(filenames, ", ")))
}
func (c *Controller) MdPage(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "fileName")
	var result bson.M

	filter := bson.D{primitive.E{Key: "slug", Value: fileName}}
	projection := options.FindOne().SetProjection(bson.D{primitive.E{Key: "content", Value: 1}})

	err := c.DB.MdCollection.
		FindOne(r.Context(), filter, projection).
		Decode(&result)
	if err != nil {
		WriteErrorResponse(w, http.StatusNotFound, notFoundErrorMsg, err)
		return
	}

	content, ok := result["content"].(string)
	if !ok {
		WriteErrorResponse(w, http.StatusInternalServerError, internalServerErrorMsg, err)
		return
	}

	body, toc := lib.MdParse(content)
	render := components.MdPage(fileName, body, toc)
	hxRequest := r.Header.Get("HX-Request")
	if hxRequest == "true" {
		lib.Htmx(w, r, render)
	} else {
		lib.Htmx(w, r, layouts.Main(render))
	}
}
