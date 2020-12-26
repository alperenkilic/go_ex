// goroutine anlamak için daha basit bir örnek
package main

import (
	"fmt"
	"time"
)

func main() {
	// countName ve countSurname fonksiyonlarının başında "go" olmaz ise countSurname fonksiyonu asla çalışmaz
	// goroutine ile iksini de eş zamanlı çalıştırıyoruz.
	go countName("alperen")
	go countSurname("kılıç")
	//fmt.Scanln() // --> scanln sonsuza kadar değer topladığı için sonuçlar yazdırılır
	// scanln olmazsa değerler yazdırılmaz.
	time.Sleep(time.Second * 4) // 4 saniye boyunca çalışır.
}

func countName(name string) {
	for i := 1; true; i++ {
		fmt.Printf("%v  %s\n", i, name)
		time.Sleep(time.Second)
	}
}

func countSurname(surname string) {

	for i := 1; true; i++ {
		fmt.Printf("%v  %s \n", i, surname)
		time.Sleep(time.Second)
	}
}
