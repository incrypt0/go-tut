package main

import (
	"fmt"
)

func main() {

	//for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	s := [3]string{"Apple", "Orange", "Pineapple"}

	//looping through arrays
	for _, item := range s {
		fmt.Println(item)
	}
}
