package main

import "fmt"

func main() {

	name, age := "alperen", 22

	//fmt.Print("benim adim ", name, " yasim ", age) 				// boşluk moşluk koymaz dümdüz yazdırır
	//fmt.Println("benim adım", name, "ve ben", age, "yasindayım")  //--> boşlukları kendi koyar println'in
	fmt.Printf("benim adım %v ve ben %v yaşındayım", name, age)
	//fmt.Printf("benim adim %v\n", name)
	/* 	fmt.Print(name)
	fmt.Printf(name) */ // prtintFormatted. formatlıdır

}
