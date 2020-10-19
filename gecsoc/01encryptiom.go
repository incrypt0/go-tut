package main

import (
	"fmt"
)

func main() {
	var k int
	var word string
	fmt.Scanln(&k)
	fmt.Scanln(&word)
	runes := []rune(word)

	for i := 1; i <= k; i++ {
		item := runes[:i]
		itemPlus := runes[i:]
		fmt.Println(string(item), string(itemPlus))
		runes = append([]rune{}, append(itemPlus, item...)...)
		// fmt.Println(string(runes))
	}
	fmt.Println(string(runes))
}
