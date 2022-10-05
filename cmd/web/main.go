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

	/* indexfor all web pages*/
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
