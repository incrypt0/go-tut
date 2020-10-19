package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("Enter the string :")
	fmt.Scan(&x)
	if strings.Contains(x, "a") &&
		strings.HasPrefix(x, "i") &&
		strings.HasSuffix(x, "n") {
		fmt.Println("Found!!")
	} else {
		fmt.Println("Not Found!!")
	}

}
