package main

import "fmt"

/* func main() {

	var employee struct {
		name      string
		isMarried bool
		age       int
	}
	//fmt.Println(employee) //--> structu gösterir
	fmt.Println(employee.age) //--> empoyee struct'ının age alanını gösterir
} */
type employee struct {
	name      string
	age       int
	isMarried bool
	kids      []string
}

type engineer struct {
	employee //ENGINEER has a EMPLOYEE --- is a has a farkını anla. is a olunca aynı tür has a olunca özelliğine sahip
	skills   []string
}

type easytest struct {
	name      string
	age       int
	isMarried bool
	kids      []string
	skills    []string
}

func main() {

	var e1 employee
	e1.name = "alperen"
	e1.age = 22
	e1.isMarried = true
	e1.kids = []string{"eylül", "ekim"}

	fmt.Println(e1)

	// veya şöyle

	e2 := employee{
		name:      "alperen",
		age:       22,
		isMarried: true,
		kids:      []string{"eylül", "ekim"},
	}
	fmt.Println(e2)

	eng1 := engineer{
		employee: employee{
			name:      "alp",
			age:       22,
			isMarried: true,
			kids:      []string{"al", "pe", "ren"}},
		skills: []string{"tech", "speak"},
	}

	fmt.Println(eng1)

	et1 := easytest{
		name:      "alperen",
		age:       22,
		isMarried: true,
		kids:      []string{"al", "fon", "zo"},
		skills:    []string{"l", "e", "g", "e", "d"},
	}
	fmt.Println(et1)

	//////////--ANONIM STRUCT--/////////

	patron := struct {
		name    string
		surname string
	}{name: "mehmet", surname: "fatih"}
	fmt.Println(patron)
}
