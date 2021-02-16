package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("dosya.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// dosya açıldı

	newFile, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer newFile.Close()
	// dosya oluşturuldu

	yazılanDosya, err := io.Copy(newFile, file)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalf("%d byte yazıldı", yazılanDosya)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
