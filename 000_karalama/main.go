package main

import "fmt"

type daire struct {
	r float64
}

func (s daire) alanBul() float64 {
	return 3.14 * s.r * s.r
}

func (s daire) cevreBul() float64 {
	return 2 * 3.14 * s.r
}

type kare struct {
	kenar float64
}

func (s kare) alanBul() float64 {
	return s.kenar * s.kenar
}

func (s kare) cevreBul() float64 {
	return s.kenar * 4
}

type shape interface {
	cevreBul() float64
	alanBul() float64
}

func interfacFunc(s shape) {
	fmt.Println(s)
	fmt.Println(s.alanBul())
	fmt.Println(s.cevreBul())
}

func main() {

	a := kare{5}
	fmt.Println(a.alanBul())
	fmt.Println(a.cevreBul())
	b := daire{5}
	fmt.Println(b.alanBul())
	fmt.Println(b.cevreBul())

	fmt.Printf("\n\n\n")
	interfacFunc(a)
	fmt.Printf("\n\n")
	interfacFunc(b)
}
