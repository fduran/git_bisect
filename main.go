package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey! %s\n", r.URL.Path)
	// add logging
	fmt.Printf("Request: %s\n", r.URL.Path)
}

func main() {
	// serve on :8181
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8181", nil)
}
