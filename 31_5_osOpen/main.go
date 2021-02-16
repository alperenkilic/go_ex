package main

import (
	"log"
	"os"
)

func main() {
	/* file, err := os.Open("dosya.txt")

	if err != nil {
		log.Println(err)
	}
	defer file.Close() */ // bu şekilde önizleme gibi açılır

	file, err := os.OpenFile("dosya.txt", os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	/*os.O_APPEND=ekleme
	os.O_CREATE=oluşturma
	os.O_RDONLY=sadece okuma
	os.O_WRONLY=sadece yazma

	os.O_WRONLY|os.O_CREATE şeklinde birden fazla ayarlanabilir
	*/

}
