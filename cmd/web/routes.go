package main

import (
	"net/http"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux := http.NewServeMux()

	dynamic := alice.New(app.sessionManager.LoadAndSave)
	
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.SnippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.SnippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.SnippetCreatePost))

	standart := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standart.Then(mux)
}