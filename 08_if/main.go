package main

import "fmt"

func main() {
	s := 12
	if x := 2; x < 0 {
		fmt.Println("git testi--x negatiftir")
	} else if x%2 == 1 {
		fmt.Println("x çifttir")
	} else {
		if x == 12 {
			fmt.Println("else içindeki if 1 ")
		} else if x < 14 {
			fmt.Println("x 15 den küçük")
		}

	}
	fmt.Println(s, " sayısı ")
}
