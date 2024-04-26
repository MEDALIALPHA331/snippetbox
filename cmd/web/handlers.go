package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func HandleIndex(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Fatal(err.Error())
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HandleGetItem(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, req)
		return
	}

	fmt.Fprintf(writer, "wow the id of the view is %d", id)
}
func HandleSnippetForm(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello world"))
}

func HandlePostSnippet(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusCreated)

	writer.Write([]byte("snippet created...."))
}
