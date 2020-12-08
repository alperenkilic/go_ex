package main

import "fmt"

func main() {
	x := 22

	fmt.Println(x)     //// value
	fmt.Println(&x)    //// adress of 22
	fmt.Println(*(&x)) ///// dereference
	fmt.Println(&*&*&*&*&*&*&*&*&*&*&x)

	// & ---> shows the value's adress
	// * ---> shows the adress's value
}
