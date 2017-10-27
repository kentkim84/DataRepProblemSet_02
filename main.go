package main

import (
	"net/http"
	"log"
	"html/template"
)

type PageVariables struct {
	HeadTitle string
	BodyTitle string
	Response string
  }

func handler(w http.ResponseWriter, r *http.Request) {
	// set local variables
	headTitle := "Guessing Game"
	bodyTitle := "Guessing Game"
	response := r.URL.Path[1:]
	// initialise local variables to struct
	MyPageVariables := PageVariables {
		HeadTitle : headTitle,
		BodyTitle : bodyTitle,
		Response : response,
	}
	// parse the template
	t, err := template.ParseFiles("index.html")
	if err != nil { // if there is an error
  		log.Print("template parsing error: ", err) // log error message
	}
    t.Execute(w, MyPageVariables)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}