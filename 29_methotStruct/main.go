package main

import "fmt"

type user struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	age       int
	pay       *payment
}

type payment struct {
	salary float64
	bonus  float64
}

func newUser() *user {
	u := new(user)
	u.pay = new(payment)
	return u
}

func newPayment() *payment {
	p := new(payment)
	return p
}

func (u *user) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *user) GetBornDate() int {
	return 2020 - u.age
}

func main() {

	u1 := user{
		ID:        "yetkili",
		FirstName: "alperen",
		LastName:  "kılıç",
		UserName:  "alperenkilic",
		age:       22,
		pay: &payment{
			salary: 6500,
			bonus:  1300,
		},
	}
	fmt.Println(u1)

	fmt.Println(u1.GetFullName())
	fmt.Println(u1.GetBornDate())
	fmt.Println(newUser())
	fmt.Println(u1)

	// 2. yöntem

	var u2 user
	// veya --> u2:=newUser()

	u2.age = 22
	u2.FirstName = "haktan"
	u2.LastName = "ugur"
	u2.ID = "h995"
	u2.pay = &payment{salary: 1200, bonus: 420}
	fmt.Println(u2)
	fmt.Println(u2.GetBornDate())
	fmt.Println(u2.GetFullName())
}
