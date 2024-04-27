package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/MEDALIALPHA331/snippetbox/internal/config"
)

var cfg config.Config

func main() {
	// err := Config.ParseConfigFromEnv()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	flag.IntVar(&cfg.PORT, "port", 5000, "HTTP Port")
	flag.StringVar(&cfg.Address, "addr", ":5000", "HTTP Network adress")
	flag.StringVar(&cfg.StaticDirPath, "static", "./ui/static/", "Static files dir path")
	flag.Parse()

	

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.StaticDirPath))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", HandleIndex)
	mux.HandleFunc("GET /snippet/view/{id}", HandleGetItem)
	mux.HandleFunc("GET /snippet/create", HandleSnippetForm)
	mux.HandleFunc("POST /snippet/create", HandlePostSnippet)

	log.Printf("Server is running on Port %d", cfg.PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), mux)
	log.Fatal(err)
}
