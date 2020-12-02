package main

import "fmt"

const x = 2222

var y int16 = 1200

/* func main() {
	const age int = 24

	fmt.Printf("%T %v\n", x, x)

	fmt.Printf("%T %v\n", y, y)

	fmt.Printf("%T %v\n", x+y, x+y) // typeless x bu kısımda geçici olarak int16 formatına dönüşüyor.

	fmt.Printf("%T %v\n", x, x)

} */

func main() {
	const x float64 = 6.4
	y := 4 + x

	fmt.Printf("%T %v\n", x, x)

	fmt.Printf("%T %v\n", y, y)

}
