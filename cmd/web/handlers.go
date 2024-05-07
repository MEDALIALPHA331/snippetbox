package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *Application) HandleIndex(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(writer, req, err)
		return
	}

	err = ts.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		app.serverError(writer, req, err)
	}
}

func (app *Application) HandleGetItem(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, req)
		return
	}

	fmt.Fprintf(writer, "wow the id of the view is %d", id)
}
func (app *Application) HandleSnippetForm(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello world"))
}

func (app *Application) HandlePostSnippet(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusCreated)

	writer.Write([]byte("snippet created...."))
}
