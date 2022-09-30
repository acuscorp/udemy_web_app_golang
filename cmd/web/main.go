package main

import (
	"net/http"

	"github.com/acuscorp/go-course/pkg/handlers"
)

const PORT_NUMBER = ":8080"

func main() {
	/* indexfor all web pages*/
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
