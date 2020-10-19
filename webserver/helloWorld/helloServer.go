package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Go is Awesome %s</h1>", "Isn't It?")
}
