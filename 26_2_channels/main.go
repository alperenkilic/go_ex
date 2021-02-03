package main

import "fmt"

func topla(s []int, c chan int) {
	toplam := 0
	for _, x := range s {
		toplam += x
	}
	c <- toplam
}

func main() {
	veri := []int{1, 2, 3, 4, 5, 6, 7}
	kanal := make(chan int)

	ver2 := []int{5, 5, 5, 5}
	kanal2 := make(chan int)

	go topla(ver2, kanal2)

	go topla(veri, kanal)

	x := <-kanal
	fmt.Println(x)

	y := <-kanal2
	fmt.Println(y)
}
