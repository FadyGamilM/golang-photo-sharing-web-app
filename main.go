package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
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

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// set the headers of the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// get the template path to render it
	templatePath := filepath.Join("templates", "home.gohtml")

	// get the result of parsing
	parsedTemplate := ParseTemplate(w, templatePath)

	// now execute it
	ExecuteTemplate(w, parsedTemplate)

	// set the status code
	w.WriteHeader(http.StatusOK)
}

func contactHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// set the headers of the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// get the file path of the template we need to render
	templatePath := filepath.Join("templates", "contact.gohtml")

	// parse the template
	parsedTemplate := ParseTemplate(w, templatePath)

	// execute the template
	ExecuteTemplate(w, parsedTemplate)

	// set the status code
	w.WriteHeader(http.StatusOK)
}

func FAQHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// set the headers of the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// get the template path to render it
	templatePath := filepath.Join("templates", "faq.gohtml")

	// parse the template
	parsedTemplate := ParseTemplate(w, templatePath)

	// execute the template
	ExecuteTemplate(w, parsedTemplate)

	// set the response status code
	w.WriteHeader(http.StatusOK)
}

// our custom router which implementes the `Handler` interface
type Router struct{}

// implement the interface specs
func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandlerFunc(w, r)
		return

	case "/contact":
		contactHandlerFunc(w, r)
		return

	case "/FAQ":
		FAQHandlerFunc(w, r)
		return

	default:
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "<h1> Page Not Found ! </h1>")
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
}

func main() {
	// instantiate our custom router
	// router := Router{}

	// http.HandleFunc("/", homeHandlerFunc)
	// http.HandleFunc("/contact", contactHandlerFunc)
	// http.HandleFunc("/", pathRouterHandler)
	r := chi.NewRouter()
	r.Get("/", homeHandlerFunc)
	r.Get("/contact", contactHandlerFunc)
	r.Get("/faq", FAQHandlerFunc)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("server is running on port 3000")

	// http.ListenAndServe("127.0.0.1:3000", http.HandlerFunc(pathRouterHandler))
	http.ListenAndServe("127.0.0.1:3000", r)
}
