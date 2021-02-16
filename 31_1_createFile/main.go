package main

import (
	"log"
	"os"
)

var (
	err     error
	newFile *os.File
)

func main() {
	newFile, err = os.Create("enough.txt")
	if err != nil {
		log.Fatal(err)
	}
}
