package main

import "fmt"

func main() {
	myMap := map[string]int{

		"jeff": 1,
		"elon": 2,
		"bill": 3,
	}
	fmt.Println(myMap)
	fmt.Println()
	deger, ok := myMap["jeff"] //değeri ve olup olmadığını
	fmt.Println(deger, ok)

	_, ok2 := myMap["jeff"] //olup olmadığı
	fmt.Println(ok2)

	delete(myMap, "jeff") // jeff keyini sildik
	fmt.Println(myMap)
	fmt.Println()
	myMap["alperen"] = 12 // alperen keyini ekledik
	fmt.Println(myMap)

	for k, v := range myMap { // range ile k ve v yi yazdırmak
		fmt.Println(k, v)
	}

}
