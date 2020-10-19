package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("__Go Ecommerce API__")

	// our main router
	r := mux.NewRouter()

	// handler functions
	r.HandleFunc("/", homeRouteHandler)
	r.HandleFunc("/{route}", rootHandler2)

	// http server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func homeRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	if _, err := w.Write([]byte("hello")); err != nil {
		log.Fatal(err)
	}
}

func rootHandler2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["route"])
	if _, err := w.Write([]byte("Hi there gadiye")); err != nil {
		log.Fatal(err)
	}
}
