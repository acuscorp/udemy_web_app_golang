package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const PORT_NUMBER = ":8080"

// this is the home page
func Home(w http.ResponseWriter, r *http.Request) {

}

// this is the about page
func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	/* indexfor all web pages*/
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(PORT_NUMBER, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {

	parseTemplate, _ := template.ParseFiles(templateFolderPath() + tmpl)

	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func templateFolderPath() string {
	return "./templates/"
}
