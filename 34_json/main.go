package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var name, adress string             // creating variables
	database := make(map[string]string) // if you do not make map, you struggle go routine error.
	fmt.Print("Write a name= ")
	fmt.Scanln(&name)
	fmt.Print("Write a adress= ")
	fmt.Scanln(&adress)
	database["name"] = name
	database["adress"] = adress
	jsonDatabase, err := json.Marshal(database) // json object
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(jsonDatabase))

}
