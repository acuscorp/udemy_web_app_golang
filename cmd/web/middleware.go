package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf CSRF protection to all POST reques
func NoSruve(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// session loads and save the session on eery request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
