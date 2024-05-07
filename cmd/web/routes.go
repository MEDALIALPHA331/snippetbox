package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.config.StaticDirPath))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.HandleIndex)
	mux.HandleFunc("GET /snippet/view/{id}", app.HandleGetItem)
	mux.HandleFunc("GET /snippet/create", app.HandleSnippetForm)
	mux.HandleFunc("POST /snippet/create", app.HandlePostSnippet)

	return mux
}
