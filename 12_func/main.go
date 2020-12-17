package main

import "fmt"

var p int

func main() {

	//fmt.Println(disc(1, 0, 0))
	//tanim("erzurum", 25)
	if test(10, 2, 30) == 1 {
		fmt.Println("üç değer de çifttir.")
	}
}

func disc(a, b, c int) int {
	return a + b + c
}

func tanim(x string, y int) {
	fmt.Printf("%s şehrinin plaka kodu--> %d\n", x, y)
}

func test(x, y, z int) int {
	if x%2 == 0 && y%2 == 0 && z%2 == 0 {
		p = 1
	}
	return p
}
