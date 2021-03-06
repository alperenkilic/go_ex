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

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

func zeroval(ival int) { // makes value 0
	ival = 0
}

func zeroptr(iptr *int) { // makes adress of value 0
	*iptr = 0
}
