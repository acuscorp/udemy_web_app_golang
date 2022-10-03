package handlers

import (
	"fmt"
	"net/http"

	"github.com/acuscorp/go-course/pkg/render"
)

// this is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\"home.page.tmpl\": %v\n", "home.page.tmpl")
	render.RenderTemplate(w, "home.page.tmpl")
}

// this is the about page
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\"about.page.tmpl\": %v\n", "about.page.tmpl")
	render.RenderTemplate(w, "about.page.tmpl")
}
