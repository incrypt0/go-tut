package main

import "fmt"

func insertionSort(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		item1 := arr[i]
		j := i - 1
		for j >= 0 && item1 < arr[j] {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			j--
		}

	}
	return arr
}
func main() {
	arr := [10]int{7, 3, 8, 5, 1, 9, 2, 4, 0, 6}
	fmt.Println(insertionSort(arr))
}
