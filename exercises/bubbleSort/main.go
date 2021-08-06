package main

import "fmt"

func main() {
	unsorted := []int{8, 3, 6, 9, 2, 0, 4, 5, 7, 1}
	sorted := bubbleSort(unsorted)
	fmt.Println(sorted)
}
func bubbleSort(unsorted []int) []int {
	len := len(unsorted)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len-1; i++ {
			if unsorted[i] > unsorted[i+1] {
				unsorted[i], unsorted[i+1] = unsorted[i+1], unsorted[i]
				swapped = true
			}
		}
	}
	return unsorted
}
