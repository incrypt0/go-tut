package main

import (
	"fmt"
)

func main() {

	var arr3 = [5]int{1, 2, 3, 4, 5}

	//Demonstration of range
	for a, b := range arr3 {
		fmt.Println(a, b)
	}
	fmt.Println(arr3)

}
