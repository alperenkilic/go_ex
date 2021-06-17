package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type site struct {
	XMLName     xml.Name `xml:"site"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"`
	Category    string   `xml:"Category"`
}

type ObjectSites struct {
	XMLName     xml.Name `xml:"sites"`
	Version     string   `xml:"version,attr"`
	Sites       []site   `xml:"site"`
	Description string   `xml:",innerxml"`
}

func main() {
	file, err := os.Open("sites.xml")
	defer file.Close()
	defer log.Printf("%s dosyası kapatıldı", file.Name())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s dosyası açıldı", file.Name())

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s dosyası okunuyor...", data)

	// bu ders-dosya devam edecek...

}
