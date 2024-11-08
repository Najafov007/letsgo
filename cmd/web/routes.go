package main

import (
	"net/http"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux := http.NewServeMux()
	
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.SnippetView)
	mux.HandleFunc("GET /snippet/create", app.SnippetCreate)
	mux.HandleFunc("POST /snippet/create", app.SnippetCreatePost)

	standart := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standart.Then(mux)
}