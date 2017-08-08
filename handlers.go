package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// Index : Basic web server that responds to any requests by simply outputting the request url
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, index Endpoint %q", html.EscapeString(r.URL.Path))
}

func Getall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Get All Endpoint %q", html.EscapeString(r.URL.Path))

	todos := Todos{
		Todo{Name: "Write presentation", Completed: false},
		Todo{Name: "Host meetup", Completed: true},
	}

	json.NewEncoder(w).Encode(todos)

}

func Getid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Get Id Endpoint %q", html.EscapeString(r.URL.Path))
}
func Createid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Post Id Endpoint %q", html.EscapeString(r.URL.Path))
}

func Deleteid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Delete Id Endpoint %q", html.EscapeString(r.URL.Path))
}
