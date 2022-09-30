package main

import (
	"net/http"
)

const PORT_NUMBER = ":8080"

func main() {
	/* indexfor all web pages*/
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
