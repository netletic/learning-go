package main

import (
	"net/http"

	"github.com/justinas/alice"
	"snippetbox.netletic.com/ui"
)

func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()

	// static responses
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// dynamic responses | unauthenticated paths
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	// dynamic responses | authenticated paths
	protected := dynamic.Append(app.requireAuthentication)
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	// standard middleware for all requests
	standard := alice.New(
		app.injectTracing,
		app.logHTTPExchange,
		app.recoverPanic,
		injectSecurityHeaders,
	)
	return standard.Then(mux)
}
