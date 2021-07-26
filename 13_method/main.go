package main

import "fmt"

type dogs struct {
	name string
	age  int
}

func (dog dogs) list() {
	fmt.Println(dog.name)
	fmt.Println(dog.age)
}

func main() {
	myDog := dogs{"nut", 4}
	myDog.list()
}
