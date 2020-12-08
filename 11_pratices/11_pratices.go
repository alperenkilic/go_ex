package main

import "fmt"

/* func main() {
	for x := 0; x <= 10; x++ {
		if x%2 == 0 {
			fmt.Println(x, " çifttir")
		} else {
			fmt.Println(x, " tektir")
		}
	}
} */

/* func main() {
	x := 1
	if x%2 == 0 {
		fmt.Println(x, " çifttir")
		return
	}
	fmt.Println(x, " tektir")
} */

func main() { // bu kod neden çalışmıyor ??s
	var x int
	var y int
	for x = 0; x < 20; x++ {
		for y = 2; y < x-1; y++ {
			if x%y == 0 {
				fmt.Println(x, "asal değil")
			}
			break
		}
	}

}

//var counter int

/* func main() { ///////////// asal hesaplar.
	for i := 2; i < 20; i++ {
		for y := 2; y < i; y++ {
			if i%y == 0 {
				counter++
			}
		}
		if counter == 0 {
			fmt.Println(i, "asldır")
		}
		counter = 0
	}
} */
