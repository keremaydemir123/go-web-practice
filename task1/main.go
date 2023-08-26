package main

import (
	"fmt"
	"net/http"
)

type router int

var PORT = ":8080"

func (r router) handleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Server runs on port: %s", PORT)
}

func (r router) handleDog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is dog route")
}

func (r router) handleMe(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello my name is Kerem")
}

func main() {
	var r router

	http.HandleFunc("/", r.handleRoot)
	http.HandleFunc("/dog", r.handleDog)
	http.HandleFunc("/me", r.handleMe)

	http.ListenAndServe(PORT, nil)
}
