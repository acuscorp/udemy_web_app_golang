package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmplName string) {
	// crerate a template cache
	tmplCache, err := createTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tmplCache[tmplName]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
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

func templatesFolderPath() string {
	return "./templates/"
}
