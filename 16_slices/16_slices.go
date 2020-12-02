package main

import "fmt"

func main() {

	myArr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(myArr, "array1")
	myArr2 := myArr
	fmt.Println(myArr2, "array2")
	myArr2[0] = 999
	fmt.Println(myArr2, "array2'nin 0 ı değişti")
	fmt.Println(myArr, "array1 değişmedi")
	fmt.Println()
	////////////////////////////////////////////////////////////
	/////	SLICES	//////////////////////////////////////////
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice, "slice1")
	mySlice2 := mySlice
	fmt.Println(mySlice2, "slice2")
	mySlice2[0] = 11111
	fmt.Println(mySlice2, "slice2'nin 0 ı değişti")
	fmt.Println(mySlice, "slice1'in değeri değişti ")

	mySlice3 := make([]int, 5) // make ile slice oluşturmaca<<<<<<<<<<<<<<<<<--------------
	fmt.Println(mySlice3)

	mySlice4 := myArr[1:4] // myArr'in 1'inci elemanından 4.elemanına. 4 dahil değil.
	fmt.Println(mySlice4)

	mySlice4 = append(mySlice4, 7, 8, 9) //7,8,9 u ekledik
	fmt.Println(mySlice4)

	mySlice5 := append(mySlice4, mySlice3...) // slice'ın içine diğer slice'ın değerleri.
	fmt.Println(mySlice5)
	fmt.Println()
	//eleman silme
	mySlice5 = mySlice5[2:len(mySlice5)]
	fmt.Println(mySlice5)

	mySlice5 = mySlice5[:len(mySlice5)-2]
	fmt.Println(mySlice5)

	var mySlc []int //== nill slice
	fmt.Printf("%#v \n", mySlc)

	mySlcc := make([]int, 0) //==boş slice
	fmt.Println(mySlcc)
}
