package main

import "fmt"

// Anonim + variadic fonksiyon

func main() {

	toplaSöyle := func(sayılar ...int) (toplam, uzunluk int) {
		// dönen değerler fonksiyon tanımında yazıldığı için çıplak return yapıldı.
		for _, i := range sayılar {
			toplam += i
			uzunluk = len(sayılar)
		}
		return
	}

	total, leng := toplaSöyle(2, 3, 5, 6)
	fmt.Println(total, leng)
}
