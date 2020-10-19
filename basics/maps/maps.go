package main

import "fmt"

func main() {
	var idMap1 map[string]int
	idMap1 = make(map[string]int)
	idMap2 := map[string]int{
		"Raj":   1,
		"Varun": 2,
	}
	fmt.Println(idMap1, idMap2)
	fmt.Println(idMap2["RollNo"])

	//Iterating through a map

	for a, b := range idMap2 {
		fmt.Println(a, b)
	}
}
