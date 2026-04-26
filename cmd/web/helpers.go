package main

import (
	"html/template"
	"net/http"
)

// Render currently renders the individual page specified from the handler call
// however it is meant to support templates.go newTemplateCache() at some point.
func (app *App) render(w http.ResponseWriter, status int, page string, data *templateData) {

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/partials/login.html",
		page,
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
