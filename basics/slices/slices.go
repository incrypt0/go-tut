package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//... means its is still an array but the no of elements is automatically inferred
	slice1 := a[1:4]
	slice2 := a[2:5]
	slice3 := make([]int, 3) //specifying length manually
	slice4 := make([]int, 3, 5)
	slice4 = append(slice4, 2)
	fmt.Println(slice1, slice2, slice3, slice4)
	fmt.Println(len(slice1), cap(slice1))
}
