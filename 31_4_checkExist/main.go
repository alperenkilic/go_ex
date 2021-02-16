package main

import (
	"fmt"
	"log"
	"os"
)

var fileInfo *os.FileInfo
var err error

func main() {

	fileInfo, err := os.Stat("buradayım.txt")
	if err != nil {
		if os.IsNotExist(err) { // bu satırı tam anlamadım.
			log.Fatal(err)
		}
	} else {
		fmt.Println("dosya mevcut, bilgileri : ")
		log.Println(fileInfo)
	}

}
