package handlers

import (
	"net/http"

	"github.com/acuscorp/go-course/pkg/config"
	"github.com/acuscorp/go-course/pkg/models"
	"github.com/acuscorp/go-course/pkg/render"
)

// the repository used by the handles
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// this is the home page
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// this is the about page
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
