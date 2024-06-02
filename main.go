package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello: %s\n", r.URL.Path)
}

func main() {
	// serve on :8080
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
