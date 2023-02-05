package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/FadyGamilM/photosharing/controllers"
	"github.com/FadyGamilM/photosharing/views"
	"github.com/go-chi/chi/v5"
)

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
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
}

func main() {
	// => VERSION [4] using chi
	// => create the router instance
	r := chi.NewRouter()

	// => parse the templates before our server booted up and run
	parsingResult := views.ParseTemplate(filepath.Join("templates", "home.gohtml"))
	// => check if there is any error while parsing
	if parsingResult.Err != nil {
		panic(parsingResult.Err)
	}
	// => now we can utilize our controller to render the template
	r.Get("/", (controllers.StaticHandler(parsingResult)))

	parsingResult = views.ParseTemplate(filepath.Join("templates", "contact.gohtml"))
	if parsingResult.Err != nil {
		panic(parsingResult.Err)
	}
	r.Get("/contact", (controllers.StaticHandler(parsingResult)))

	parsingResult = views.ParseTemplate(filepath.Join("templates", "faq.gohtml"))
	if parsingResult.Err != nil {
		panic(parsingResult.Err)
	}
	r.Get("/faq", (controllers.StaticHandler(parsingResult)))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("server is running on port 3000")

	http.ListenAndServe("127.0.0.1:3000", r)

	// => VERSION [1] of regstiering a router handler
	// http.ListenAndServe("127.0.0.1:3000", http.HandlerFunc(pathRouterHandler))

	// => VERSION [2] of registering a router handler
	// instantiate our custom router
	// router := Router{}
	// http.ListenAndServe("localhost:3000", router)

	// => VERSION [3] of registering a router handler
	// http.HandleFunc("/", homeHandlerFunc)
	// http.HandleFunc("/contact", contactHandlerFunc)
	// http.HandleFunc("/", pathRouterHandler)
}
