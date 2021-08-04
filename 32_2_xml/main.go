package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type jsonPerson struct {
	FirstName string
	LastName  string
	Title     string
	Division  string
	Building  int
	Room      string
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	Title     string   `xml:"title"`
	Division  string   `xml:"division"`
	Building  int      `xml:"building"`
	Room      string   `xml:"room"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

func (p Person) String() string {
	return fmt.Sprintf("FirstName: %s\t\tLastName: %s\t\tTitle: %s\t\tDivision: %s\t\tBuilding: %d\t\tRoom: %s\n", p.FirstName, p.LastName, p.Title, p.Division, p.Building, p.Room)
}

func main() {
	xmlFile, err := os.Open("company.xml")
	if err != nil {
		log.Println(err)
		return
	}
	defer xmlFile.Close()
	xmlData, _ := ioutil.ReadAll(xmlFile)
	var c Company
	xml.Unmarshal(xmlData, &c)
	// fmt.Println(c.Persons) shows xml format
	// converting json

	var person jsonPerson
	var persons []jsonPerson

	for _, v := range c.Persons {
		person.FirstName = v.FirstName
		person.LastName = v.LastName
		person.Title = v.Title
		person.Division = v.Division
		person.Building = v.Building
		person.Room = v.Room

		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
	jsonFile, err := os.Create("company.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
}
