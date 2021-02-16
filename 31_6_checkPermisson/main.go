package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("dosya.txt", os.O_RDWR, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("okuma reddedildi")
		}
	}
	defer file.Close()
}
