package main

import "fmt"

/* func main() {

	for {
		fmt.Println("loop")
		break
	}
} */

////////////////// REVERSE WRITE INT //////////////////

/* unc main() {
	tersSayi := []int{}
	var sayi int
	fmt.Print("3 basamaklı sayı giriniz = ")
	fmt.Scanln(&sayi) //yazi değişkenine değer girilmesini istedik.
	fmt.Println("Girdiğiniz sayı = ", sayi)

	sonRakam := sayi % 10
	tersSayi = append(tersSayi, sonRakam)
	sayi /= 10

	ortaRakam := sayi % 10
	tersSayi = append(tersSayi, ortaRakam)
	sayi /= 10

	ilkRakam := sayi % 10
	tersSayi = append(tersSayi, ilkRakam)
	sayi /= 10

	sonuc := tersSayi[0]*100 + tersSayi[1]*10 + tersSayi[2]*1
	fmt.Println(sonuc)
} */

////////////////// REVERSE WRITE INT ///////////////////
// git test
//////////////////  INTERTWINDED ARRAYS ///////////////////

/* func main() {

	var twoD [5][7]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
} */

//////////////////  ///////////////////

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
