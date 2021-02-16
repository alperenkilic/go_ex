package main

import (
	"fmt"
	"log"
	"os"
)

var fileInfo *os.FileInfo
var err error

func main() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file name  ", fileInfo.Name())
	fmt.Println("file size : ", fileInfo.Size())
	fmt.Println("file permissons :", fileInfo.Mode())
	fmt.Println("isDir?? : ", fileInfo.IsDir())
}
