package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandleIndex(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("Server", "Go")
	writer.Write([]byte("Hello from snippetbox"))
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
