package main

import "fmt"

type rectangle struct {
	edge float64
}

func area(r rectangle, c chan float64) {
	result := r.edge * r.edge
	c <- result
}

func main() {

	myChan := make(chan float64)
	kare := rectangle{5}
	go area(kare, myChan)
	fmt.Println(<-myChan)
}
