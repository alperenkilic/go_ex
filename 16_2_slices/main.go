package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	storage := make([]int, 3) // if I use "make slice" go complier saying panic error
	var input string
	for {
		fmt.Printf("enter number or 'x' for quit= ")
		fmt.Scanln(&input)

		if input == "x" {
			fmt.Println("program closed")
			l := len(storage)
			for i := 0; i < l; i++ { // standart bubble sorting algorithm
				for j := 0; j < (l - 1 - i); j++ {
					if storage[j] > storage[j+1] {
						storage[j], storage[j+1] = storage[j+1], storage[j]
					}
				}
			}

			fmt.Println(storage)
			break
		}

		buffer, _ := strconv.Atoi(input)

		if i > 2 {
			storage = append(storage, buffer)
			i++
		} else {
			storage[i] = buffer
			i++
		}
	}
}
