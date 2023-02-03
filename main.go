package main

import (
	"fmt"
	"net/http"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Home Page </h1>")
}

func contactHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `
      <h1>Contact Page</h1>
      <p>For more info, you can contact with me at <a href=/"gamilfady605@gmail.com/"> gamilfady605@gmail.com </a> </p>
   `)
}

func FAQHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,
		`
		<h1> FAQ Page </h1>
		<h3> Can i have a free trial ? </h3>
			<p> Yes, we have a free trial for 30 days with money back guranteed </p>
		<h3> Do you have a supportl ? </h3>
			<p> Yes, you can contact us at support@photoSharing.com we have a support team answering your questions 24/7 </p>
		`)
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
	router := Router{}

	// http.HandleFunc("/", homeHandlerFunc)
	// http.HandleFunc("/contact", contactHandlerFunc)
	// http.HandleFunc("/", pathRouterHandler)

	fmt.Println("server is running on port 3000")

	// http.ListenAndServe("127.0.0.1:3000", http.HandlerFunc(pathRouterHandler))
	http.ListenAndServe("127.0.0.1:3000", router)
}
