package main

import "fmt"

func main() {
	var sayi int = 12
	var sayi2 float32 = 1.22

	var kocaman int16 = 20000

	fmt.Printf("%T  %T \n", sayi, sayi2)

	x := int(sayi2)
	fmt.Println(sayi + int(sayi2) + x)
	fmt.Println(int8(kocaman))

}
