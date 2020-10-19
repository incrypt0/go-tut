package main

import (
	"fmt"
	"strconv"
)

func swap(num1 *int, num2 *int) {
	*num1, *num2 = *num2, *num1
}
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	fmt.Println(arr)
}
func main() {
	var arr []int
	var a string
	fmt.Println("Enter the numbers in new line and Enter X when you are finsihed entering the numbers :")
l1:
	for i := 0; true; i++ {
		fmt.Scanf("%v", &a)

		if a == "X" || a == "x" {
			break l1
		}
		numofA, _ := strconv.Atoi(a)
		arr = append(arr, numofA)

	}
	BubbleSort(arr)
}
