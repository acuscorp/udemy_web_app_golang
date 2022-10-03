package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// this parse files, and can receive one or more files, in this case the content and de base
	parseTemplate, _ := template.ParseFiles(templatesFolderPath()+tmpl, templatesFolderPath()+"base.layout.tmpl")
	fmt.Printf("tmpl: %v\n", tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprint(templatesFolderPath(), t),
		fmt.Sprint(templatesFolderPath(), "base.layout.tmpl"),
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
}

func templatesFolderPath() string {
	return "./templates/"
}
