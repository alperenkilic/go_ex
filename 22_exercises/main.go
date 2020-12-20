package main

import "fmt"

//////////////////////////////////////////////////////////////////////
/* type myType int

 func main() {

	var x myType
	x = 10

	fmt.Printf("X'in veri tipi %T 'dir", x)
} */

/* func main() {
	x := 10
	fmt.Println(x)
	y := &x
	*y = 20
	fmt.Println(*y)
	fmt.Println(x)
} */
/////////////////////////////////////////////////////////////////////////7
/*
type rectangle struct {
	a int
	b int
}

func (dikt rectangle) area() int {
	return dikt.a * dikt.b
}

func (dikt rectangle) circumference() int {
	return 2 * (dikt.a + dikt.b)
}

func (dikt rectangle) pisa() float64 {
	x := math.Pow(float64(dikt.a), 2)
	y := math.Pow(float64(dikt.b), 2)
	return math.Sqrt(x + y)
}

func main() {
	r1 := rectangle{3, 4}
	fmt.Println(r1.area())
	fmt.Println(r1.circumference())
	fmt.Println(r1.pisa())

} */

//////////////////////////////////////////////////////////////////////

type family struct {
	name      string
	age       int
	isMarried bool
}

func main() {
	aile := showFamily()
	fmt.Println(aile)

}

func showFamily() []family {

	f1 := family{
		"alperen",
		22,
		false,
	}

	f2 := family{
		"zehra",
		21,
		false,
	}

	return []family{f1, f2}
}
