package main

import (
	"net/http"

	"github.com/jfilipedias/snippetbox/ui"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.HandleFunc("GET /ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /about", dynamic.ThenFunc(app.about))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /signup", dynamic.ThenFunc(app.signup))
	mux.Handle("POST /signup", dynamic.ThenFunc(app.signupPost))
	mux.Handle("GET /login", dynamic.ThenFunc(app.login))
	mux.Handle("POST /login", dynamic.ThenFunc(app.loginPost))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("GET /account", protected.ThenFunc(app.accountView))
	mux.Handle("POST /logout", protected.ThenFunc(app.logoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
