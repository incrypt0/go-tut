package main

import "fmt"

func main() {
	defer fmt.Println("This will get executed after everything else")
	fmt.Println("Hi")
}
