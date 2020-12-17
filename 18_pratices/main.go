package main

import "fmt"

/* func main() {
	myArr := [...]string{"a", "b", "c", "d", "e"} // array oluşturma
	mySlice := []int{1, 2, 3, 4, 5}               //slice oluşturma

	for index, value := range myArr { //for range ile yazdırma
		fmt.Println(index, value)
	}

	for index, value := range mySlice {
		fmt.Println(index, value)
	}
} */

/* func main() {

	mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstSlice := mySlice[:4]
	fmt.Println(firstSlice)
	fmt.Println() // boşluk
	secondSlice := mySlice[4:7]
	fmt.Println(secondSlice)
	fmt.Println()
	thirdSlice := mySlice[6:9]
	fmt.Println(thirdSlice)
	fourthSlice := mySlice[2:4]
	fifthSlice := mySlice[6:8]
	finalSlice := append(fourthSlice, fifthSlice...)
	fmt.Println(finalSlice)
	fmt.Println()
	noWay := append(mySlice[2:4], mySlice[6:8]...)
	fmt.Println(noWay)
} */

func main() {

	author := map[string][]string{

		"yaşar kemal":     []string{"ince mehmed", "yer demir gök bakır"},
		"john verdon":     []string{"kurt gölü", "gözlerini ssk"},
		"agatha christie": []string{"doğu expresinde", " cinayet"},
	}

	for k, v := range author {
		fmt.Println(k, v)
	}
	//fmt.Println(author["ömer seyfettin"][0])

	for k, v := range author {
		fmt.Println("yazarımız: ", k)
		fmt.Println("ve kitapları ")
		for i, value := range v {
			fmt.Println("\t", i+1, value)
		}
	}
}
