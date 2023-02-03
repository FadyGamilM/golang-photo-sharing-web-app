package views

import (
	"html/template"
	"log"
	"net/http"
)

// ! => Helpers for parsing and executing templates
type ParsedTempalte struct {
	template *template.Template
	err      error
}

func ParseTemplate(w http.ResponseWriter, templatePath string) ParsedTempalte {
	templ, err := template.ParseFiles(templatePath)
	if err != nil {
		// if we here, so our template has something wrong in it
		log.Printf("Error while parsing the template -> %v", err)
		http.Error(w, "Server Error - Error while parsing the HTML Page", http.StatusInternalServerError)
		return ParsedTempalte{
			err:      err,
			template: nil,
		}
	}
	return ParsedTempalte{
		template: templ,
		err:      nil,
	}
}

func ExecuteTemplate(w http.ResponseWriter, parsedTemplate ParsedTempalte) {
	err := parsedTemplate.template.Execute(w, nil)
	if err != nil {
		log.Printf("Error while executing the template -> %v", err)
		http.Error(w, "Server Error - Error while executing the HTML Page", http.StatusInternalServerError)
		return
	}
}
