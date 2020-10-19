package main

import "fmt"

func change(a *[7]int) {
	(*a)[0] = 4
}

func main() {
	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	change(&a)
	fmt.Println(a)
}
