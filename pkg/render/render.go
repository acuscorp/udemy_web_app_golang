package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/acuscorp/go-course/pkg/config"
	"github.com/acuscorp/go-course/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}
func RenderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	// check if shoudl use cache memory
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get templateCache frrom app config

	// get requested template from cache
	t, ok := tc[tmplName]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCahce := map[string]*template.Template{}

	// get all of the files names *.pate.tmpl frrom ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		fileName := filepath.Base(page)
		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCahce, err
		}
		matches, err := filepath.Glob("./templates/*layout.tmpl")

		if err != nil {
			return myCahce, err
		}
		// parse template page with layout
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCahce, err
			}
		}
		// adding the template to the chache
		myCahce[fileName] = ts
	}
	return myCahce, nil

}
