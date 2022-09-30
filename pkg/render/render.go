package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parseTemplate, _ := template.ParseFiles(templateFolderPath() + tmpl)

	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func templateFolderPath() string {
	return "./templates/"
}
