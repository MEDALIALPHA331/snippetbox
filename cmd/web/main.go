package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/MEDALIALPHA331/snippetbox/internal/config"
)

type Application struct {
	logger *slog.Logger
	config config.Config
}

var cfg config.Config

func main() {
	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(loggerHandler)

	app := Application{
		logger,
		cfg,
	}

	err := cfg.ParseConfigFromEnv()
	if err != nil {
		logger.Error(err.Error())
	}

	flag.IntVar(&cfg.PORT, "port", 5000, "HTTP Port")
	flag.StringVar(&cfg.Address, "addr", ":5000", "HTTP Network adress")
	flag.StringVar(&cfg.StaticDirPath, "static", "./ui/static/", "Static files dir path")
	flag.Parse()

	mux := app.routes()

	logger.Info("Server is Starting", "Port", cfg.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), mux)
	logger.Error(err.Error())
	os.Exit(1)
}
