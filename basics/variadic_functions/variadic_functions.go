package main

import "fmt"

func variadicFunction(a ...int){
	for _,i :=range a {
		fmt.Println(i)
	}
}
func main() {
	variadicFunction(1,2,3,4,5)
}