package main

import (
	"log"
	"net/http"

	"github.com/acuscorp/go-course/pkg/config"
	"github.com/acuscorp/go-course/pkg/handlers"
	"github.com/acuscorp/go-course/pkg/render"
)

const PORT_NUMBER = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
