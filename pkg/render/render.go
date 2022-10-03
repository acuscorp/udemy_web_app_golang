package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	// this parse files, and can receive one or more files, in this case the content and de base
	parseTemplate, _ := template.ParseFiles(templatesFolderPath()+tmpl, templatesFolderPath()+"base.layout.tmpl")
	fmt.Printf("tmpl: %v\n", tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {

	var err error

	_, inMap := tc[t]

	if !inMap {
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached template")
	}

	err = tc[t].Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}

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
	return nil
}

func templatesFolderPath() string {
	return "./templates/"
}
