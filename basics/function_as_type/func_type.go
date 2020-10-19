package main

import "fmt"

var funcVar func(int) int

func printInt(a int) int {
	return a + 1
}
func main() {
	var a int = 2
	funcVar = printInt
	fmt.Println(funcVar(a))
}
