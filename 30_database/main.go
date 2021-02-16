package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "alperen:123123/test")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	createStatement := "'users'('ID' int(11) NOT NULL AUTO_INCREMENT,'Username' varchar(45) NOT NULL, 'Email' varchar(45) NOT NULL, 'Password' varchar(45) NOT NULL,'FirstName' varchar(45) NOT NULL,'LastName' varchar(45) NOT NULL, 'BirthDate' varchar(45) DEFAULT NULL,'IsActive' tinyint(1) DEFAULT NULL,'PRIMARY KEY' ('ID')"

	_, err = db.Exec("created table if not exist" + createStatement)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("insert into users (Username,Email,Password,Firstname,Lastname,Birthdate,IsActive) VALUES('deneme','12345','test.gmail.com')")
	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserterd %d rows", rowCount)

	rows, err := db.Query("SELECT*FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

	}
}
