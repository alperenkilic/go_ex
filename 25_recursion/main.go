package main

import "fmt"

func main() {

	fmt.Println(fact(5))
	fmt.Println(topla(15))
}

// Recursion fonksiyonları destekler.

func fact(i int) int {

	if i == 0 {
		return 1
	}
	return i * fact(i-1)

}

// 0'dan girilen değere kadar olan doğal sayıları toplar.
func topla(x int) int {
	if x == 0 {
		return 0
	}
	return x + topla(x-1)
}
