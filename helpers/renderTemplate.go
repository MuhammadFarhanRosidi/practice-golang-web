package helpers

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	// Parse the template file
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
