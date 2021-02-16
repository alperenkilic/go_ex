package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("demo.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	veri := []byte("yazılacak veri satırı")

	written, err := file.Write(veri)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("yazılan veri büyüklüğü = %d", written)

	// veya

	err = ioutil.WriteFile("demo.txt", []byte("test verisi"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
