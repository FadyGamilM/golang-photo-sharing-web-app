package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/FadyGamilM/photosharing/views"
	"github.com/go-chi/chi/v5"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// set the headers of the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// get the template path to render it
	templatePath := filepath.Join("templates", "home.gohtml")

	// get the result of parsing
	parsedTemplate := views.ParseTemplate(w, templatePath)

	// now execute it
	views.ExecuteTemplate(w, parsedTemplate)

	// set the status code
	w.WriteHeader(http.StatusOK)
}

func contactHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// set the headers of the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// get the file path of the template we need to render
	templatePath := filepath.Join("templates", "contact.gohtml")

	// parse the template
	parsedTemplate := views.ParseTemplate(w, templatePath)

	// execute the template
	views.ExecuteTemplate(w, parsedTemplate)

	// set the status code
	w.WriteHeader(http.StatusOK)
}

func FAQHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// set the headers of the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// get the template path to render it
	templatePath := filepath.Join("templates", "faq.gohtml")

	// parse the template
	parsedTemplate := views.ParseTemplate(w, templatePath)

	// execute the template
	views.ExecuteTemplate(w, parsedTemplate)

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
