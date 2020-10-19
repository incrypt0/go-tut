package main

import "fmt"

func main() {
	//neew function creates a variable
	//and then returns a pointer to that variable
	a := new(int)
	*a = 2
	fmt.Println(*a)
}
