package main

import "fmt"

// interface
type shape interface {
	area() float64
}

// interface func/methot
func printArea(s ...shape) {
	for _, shape := range s {
		fmt.Println("alan: ", +shape.area())
	}
}

// structs and methots
type triange struct {
	h float64
	a float64
}

func (t triange) area() float64 {
	return t.a * t.h / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

func main() {

	kindShow := square{2}
	printArea(kindShow)
	///////////////////////////////////
	test := square{a: 4}
	printArea(test)
	///////////////////////////////////
	var kare square
	kare.a = 12
	// variadic 
	printArea(kare, kare, kare)
}
