package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// RenderTemplate is responsible for rendering data object into template file.
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

// GenerateUUID generates an UUID for in specific format.
func GenerateUUID() string {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Println("Encountered the following error when attempting to generate an UUID: ", err)
		return ""
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
