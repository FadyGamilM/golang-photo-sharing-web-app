package views

import (
	"html/template"
	"log"
	"net/http"
)

// ! => Helpers for parsing and executing templates
type Template struct {
	Template *template.Template
	Err      error
}

func ParseTemplate(templatePath string) Template {
	templ, err := template.ParseFiles(templatePath)
	if err != nil {
		// if we here, so our template has something wrong in it
		log.Printf("Error while parsing the template -> %v", err)
		return Template{
			Err:      err,
			Template: nil,
		}
	}
	return Template{
		Template: templ,
		Err:      nil,
	}
}

func (t Template) Render(w http.ResponseWriter, data interface{}) Template {
	// execute the template and send the data
	err := t.Template.Execute(w, data)
	// check if there is any error
	if err != nil {
		log.Printf("Error while executing the template -> %v", err)
		return Template{
			Template: nil,
			Err:      err,
		}
	}
	return Template{
		Err:      nil,
		Template: t.Template,
	}
}
