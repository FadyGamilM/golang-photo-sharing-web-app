package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	temp, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "fady gamil",
	}

	// take a template and process it
	err = temp.Execute(os.Stdout, user)
}
