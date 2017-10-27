package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Guessing Game"
	response := r.URL.Path[1:]
	fmt.Fprintf(w, "<h1>%s</h1>", title)
	fmt.Fprintf(w, "<p>Your guess is %s!<p>", response)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}