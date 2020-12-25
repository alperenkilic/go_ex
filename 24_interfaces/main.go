package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	a, b float64
}

func (dikt rectangle) area() float64 {
	return dikt.a * dikt.b
}

func (dikt rectangle) circumference() float64 {
	return 2 * (dikt.a + dikt.b)
}

type shape interface { // interface SADECE yapılacak fonksiyonları belirtir (signature)
	area() float64
	circumference() float64
}

type circle struct {
	r float64
}

func (daire circle) area() float64 {
	return math.Pi * (daire.r * daire.r)
}

func (daire circle) circumference() float64 {
	return 2 * math.Pi * daire.r
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference())
	fmt.Printf("%T", i)
	fmt.Println()
}

func main() {
	/* 	r1 := rectangle{3, 4}
	   	fmt.Println(r1.area())
	   	fmt.Println(r1.circumference())
	   	fmt.Println()
	   	interfaceFunc(r1) */

	cember := circle{5}
	interfaceFunc(cember) // circle'ın area-circumfence fonksiyonlarını getirir
	fmt.Println()
	r1 := rectangle{3, 4}
	interfaceFunc(r1) // rectangle'ın area-circumfence fonksiyonlarını getirir
}
