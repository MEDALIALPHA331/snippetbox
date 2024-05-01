package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/MEDALIALPHA331/snippetbox/internal/config"
)

var cfg config.Config

type Application struct {
	logger *slog.Logger
}

func main() {
	// err := cfg.ParseConfigFromEnv()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	flag.IntVar(&cfg.PORT, "port", 5000, "HTTP Port")
	flag.StringVar(&cfg.Address, "addr", ":5000", "HTTP Network adress")
	flag.StringVar(&cfg.StaticDirPath, "static", "./ui/static/", "Static files dir path")
	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(loggerHandler)

	app := Application{
		logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.StaticDirPath))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.HandleIndex)
	mux.HandleFunc("GET /snippet/view/{id}", app.HandleGetItem)
	mux.HandleFunc("GET /snippet/create", app.HandleSnippetForm)
	mux.HandleFunc("POST /snippet/create", app.HandlePostSnippet)

	logger.Info("Server is Starting", "Port", cfg.PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), mux)
	logger.Error(err.Error())
	os.Exit(1)
}
