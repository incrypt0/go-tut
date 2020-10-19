package main

import "fmt"

func lowestIndex(arr []int) int {
	var length, minIndex int = len(arr), 0
	for i := 0; i < length; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}
func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		low := lowestIndex(arr[i+1:]) + i + 1
		arr[i], arr[low] = arr[low], arr[i]
	}
	return arr
}
func main() {
	arr := []int{7, 3, 8, 5, 1, 9, 2, 4, 0, 6}
	result := selectionSort(arr)
	fmt.Println(result)
}
