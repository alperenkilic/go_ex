package main

import "fmt"

/* func main() {

	for {
		fmt.Println("loop")
		break
	}
} */

////////////////// REVERSE WRITE INT //////////////////

func main() {
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
}

////////////////// REVERSE WRITE INT //////////////////
