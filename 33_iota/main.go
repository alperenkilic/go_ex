package main

import "fmt"

const (
	a int = iota
	b
	c
	d
	e
	f
	g
)

const ()

func main() {
	fmt.Println(a, b, c, d, e, f, g)
}
