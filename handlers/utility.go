package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate is responsible for rendering data object into template file.
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}
