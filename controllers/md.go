package controllers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func Md(w http.ResponseWriter, r *http.Request) {
	dir := "./data"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
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
