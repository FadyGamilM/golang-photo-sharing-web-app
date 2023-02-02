package main

import (
	"fmt"
	"net/http"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get request is done")
	fmt.Fprintf(w, "<h1>Home Page </h1>")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func contactHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
      <h1>Contact Page</h1>
      <p>For more info, you can contact with me at <a href=/"gamilfady605@gmail.com/"> gamilfady605@gmail.com </a> </p>
   `)
}

func pathRouterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandlerFunc(w, r)
		return

	case "/contact":
		contactHandlerFunc(w, r)
		return

	default:
		fmt.Fprint(w, "<h1> 404 Not Found Page ! </h1>")
		return
	}
}

func main() {
	// http.HandleFunc("/", homeHandlerFunc)
	// http.HandleFunc("/contact", contactHandlerFunc)
	http.HandleFunc("/", pathRouterHandler)
	fmt.Println("server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
