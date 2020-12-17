package main

import "fmt"

//const x = 2222 // typeless

//var y int16 = 1200

/* func main() {
	const age int = 24

	fmt.Printf("%T %v\n", x, x)

	fmt.Printf("%T %v\n", y, y)

	fmt.Printf("%T %v\n", x+y, x+y) // typeless x bu kısımda geçici olarak int16 formatına dönüşüyor.

	fmt.Printf("%T %v\n", x, x)

} */

/* func main() {
	const x = 6.4
	y := 4 + x

	fmt.Printf("%T %v\n", x, x)

	fmt.Printf("%T %v\n", y, y)

} */

const x = 6

var y int = 12
var z float32 = 10.4

func main() {
	y = y + x
	fmt.Printf("%T  %v \n", x, x)
	fmt.Printf("%T  %v \n", y, y)

	z = z + x
	fmt.Printf("%T %v \n", z, z)
	fmt.Printf("%T %v \n", x, x)

}
