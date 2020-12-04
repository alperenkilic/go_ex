package main

import "fmt"

// how can we reverse write the array
func main() {

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	revArr := [len(arr)]int{}

	for i := range arr {
		revArr[len(arr)-1-i] = arr[i]
	}

	fmt.Println(revArr)
}
