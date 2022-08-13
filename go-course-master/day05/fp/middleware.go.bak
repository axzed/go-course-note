package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8848", http.HandlerFunc(hello))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, http: %s", r.URL.Path)
}
