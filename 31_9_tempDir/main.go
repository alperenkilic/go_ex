package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	geciciKlasor, err := ioutil.TempDir("", "belgeler")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("geçici klasör oluşturuldu: ", geciciKlasor)

	geciciDosya, err := ioutil.TempFile(geciciKlasor, "test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal("dosya oluşturuldu : ", geciciDosya.Name())

	err = geciciDosya.Close()
	if err != nil {
		log.Fatal(err)
	}

	// temizlik

	err = os.Remove(geciciDosya.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(geciciKlasor)
	if err != nil {
		log.Fatal(err)
	}

}
