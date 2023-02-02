package main

import (
	"fmt"
	"net/http"
)

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Photo Sharing App</h1>")
}

func main() {
	http.HandleFunc("/", handlerfunc)
	fmt.Println("server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
 