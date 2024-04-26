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

	homePagePath := "./ui/html/pages/home.tmpl.html"
	ts, err := template.ParseFiles(homePagePath)

	if err != nil {
		log.Fatal(err.Error())
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	ts.Execute(writer, nil)
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
