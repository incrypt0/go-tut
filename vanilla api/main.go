package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("_________Hello Go_________")

	http.HandleFunc("/", Index)
	http.HandleFunc("/register", Register)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("________ Index Handler _________")
	print(r.Method)
	_, _ = w.Write([]byte("Hello there"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register User Route")

}
