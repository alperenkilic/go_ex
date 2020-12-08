package main

import "fmt"

// look at 19_struct.go
type kilometer float32
type mile float32

// own type ile kodu daha anlaşılır kılabiliriz.
func main() {
	km := kilometer(12)
	mm := mile(33)
	fmt.Println(toKilometer(mm))
	fmt.Println(toMile(km))
}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {
	return mile(k / 1.6)
}
