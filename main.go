package main

import (
	"errors"
	"fmt"
	"net/http"
)

const PORT_NUMBER = ":8080"

// this is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page!")
}

// this is the about page
func About(w http.ResponseWriter, r *http.Request) {
	x := 5
	y := 2
	sum, _ := AddValues(x, y)
	fmt.Fprintf(w, "This is the about page and the sum of %d and %d is %d", x, y, sum)

}

func AddValues(x, y int) (int, error) {
	return x + y, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := float32(1000)
	y := float32(0)
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "%f cannot devide by 0", x)
	} else {

		fmt.Fprintf(w, "%f divided by %f is %f", x, y, f)
	}

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
