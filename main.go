package main

import (
	"errors"
	"net/http"
)

const PORT_NUMBER = ":8080"

// this is the home page
func Home(w http.ResponseWriter, r *http.Request) {

}

// this is the about page
func About(w http.ResponseWriter, r *http.Request) {

}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func main() {
	/* indexfor all web pages*/
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
