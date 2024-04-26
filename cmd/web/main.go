package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MEDALIALPHA331/snippetbox/internal/config"
)

var Config config.Config

func main() {
	err := Config.ParseConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", HandleIndex)
	mux.HandleFunc("GET /snippet/view/{id}", HandleGetItem)
	mux.HandleFunc("GET /snippet/create", HandleSnippetForm)
	mux.HandleFunc("POST /snippet/create", HandlePostSnippet)

	log.Printf("Server is running on Port %d", Config.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", Config.PORT), mux)
	log.Fatal(err)
}
