package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var hd hotdog
	err := http.ListenAndServe(":8080", hd)
	if err != nil {
		log.Fatal(err)
	}
}
