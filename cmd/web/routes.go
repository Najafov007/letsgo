package main

import (
	"net/http"

	"github.com/justinas/alice"
	"snippetbox.nijat.net/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// Unprotected route
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticated)

	// Route for testing
	mux.HandleFunc("GET /ping", ping)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.SnippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	// Protected route
	protected := dynamic.Append(app.requireAuthentication)
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.SnippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.SnippetCreatePost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	standart := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standart.Then(mux)
}