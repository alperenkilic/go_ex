package main

import "fmt"

func main() {

	//myAr := []int{1, 2, 3, 4, 3, 6, 76, 8, 87, 65, 4}
	/* cities := []string{"erzurum", "ankara", "istanbul", "izmir"}
	 	fmt.Println(myAr)
	   	fmt.Println(cities)
	   	fmt.Println(len(myAr))
	   	fmt.Println(len(cities))

	myAr[0] = (11)
	fmt.Println(myAr) */

	myAr := []int{1, 2, 3, 4, 3, 6, 76, 8, 87, 65, 4}

	/* for i := 0; i < len(myAr); i++ {
		fmt.Println(i, myAr[i])
	} */

	for i, myAr := range myAr { //for _,myAr := range myAr { <<<<----- bu şekilde sırayı yazdırmayız
		fmt.Println(i, myAr)
	}
}
