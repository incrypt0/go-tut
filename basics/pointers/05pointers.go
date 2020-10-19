package main

import (
	"fmt"
)

func main() {
	a := 12
	b := &a         //This is actually a memmory address
	fmt.Println(*b) //Prints out the memmory address
}
