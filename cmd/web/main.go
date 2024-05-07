package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/MEDALIALPHA331/snippetbox/internal/config"
	_ "github.com/go-sql-driver/mysql"
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

	db, err := startDb()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to db: %v", err.Error()))
	}
	defer db.Close()

	mux := app.routes()

	logger.Info("Server is Starting", "Port", cfg.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), mux)
	logger.Error(err.Error())
	os.Exit(1)
}

func startDb() (*sql.DB, error) {
	var conString = "root:snippetboxpwd@/snippetsdb"
	db, err := sql.Open("mysql", conString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
