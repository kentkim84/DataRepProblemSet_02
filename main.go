package main

import (
	//"fmt"
	"net/http"
	"log"
	"html/template"
)

type PageVariables struct {
	HeadTitle string
	BodyTitle string
	Response string
	Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// set local variables
	headTitle := "Guessing Game"
	bodyTitle := "Guessing Game"

	// initialise local variables to struct
	MyPageVariables := PageVariables {
		HeadTitle : headTitle,
		BodyTitle : bodyTitle,
	}
	// parse the template
	t, err := template.ParseFiles("index.html")
	if err != nil { // if there is an error
  		log.Print("template parsing error: ", err) // log error message
	}
	t.Execute(w, MyPageVariables)
	// fmt.Fprintf(w, "Execute indexHandler")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// set local variables
	headTitle := "Guessing Game"
	message := "Guess a number between 1 and 20"

	// initialise local variables to struct
	MyPageVariables := PageVariables {
		HeadTitle : headTitle,
		Message : message,
	}
	// parse the template
	t, err := template.ParseFiles("guess.tmpl")
	if err != nil { // if there is an error
  		log.Print("template parsing error: ", err) // log error message
	}
	t.Execute(w, MyPageVariables)
	// fmt.Fprintf(w, "Execute guessHandler")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/guess/", guessHandler)
    http.ListenAndServe(":8080", nil)
}