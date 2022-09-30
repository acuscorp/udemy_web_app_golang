package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// this is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// this is the about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
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
