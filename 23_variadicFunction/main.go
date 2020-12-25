package main

import "fmt"

func main() {

	fmt.Println(topla(1, 2, 3, 4, 5, 1, 2, 3, 2))

}

func topla(girilen ...int) int { // istenildiği kadar değişken girilebilir
	//variadic parametre” son parametre olarak yazılmalı ve
	// kendisinden önce türü belirtilmiş en az bir parametre değişkeni olmalı.
	toplam := 0
	for _, i := range girilen { // range yapısı kolay kullanılır
		toplam += i
	}
	return toplam
}
