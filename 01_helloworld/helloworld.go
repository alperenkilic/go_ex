package main

import "fmt"

func main() {

	/*var (
		name    string  = "haktan"
		surname string  = "kilic"
		age     int     = 22
		deger   float32 = 112.213
	)
	memleket := "erzurum" */

	var name, surname, age, deger, memleket = "alperen", "kilic", "22", "99999", "erzurum"

	/* en kısa declare yöntemi şudur */

	user1, user2, user3, loser4, plaka := "haktan", "eren", "kilic", "never", 25

	fmt.Println(user1, user2, user3, loser4, plaka)
	fmt.Println(name, surname, age, deger, memleket)
	fmt.Println(name)
	fmt.Println(memleket)
}
