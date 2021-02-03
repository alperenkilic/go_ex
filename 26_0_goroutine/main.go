package main

import (
	"fmt"
	"time"
)

// Goroutines kod hiyerarşisinden bağımsız eşzamanlı işlem yapmaya yarar.
// Bağımsız çalışmasını istediğimiz fonksiyonun başına "go" eklemek yeterlidir.
func main() {
	kahramanlar := []string{"Marvel", "Flash", "Thanos", "Flash", "Hulk", "Thor",
		"Marvel", "Flash", "Thanos", "Flash", "Hulk", "Thor"}
	go bulucuB(kahramanlar)
	go bulucuA(kahramanlar)
	go hello()
	go func() {
		fmt.Println("Merhaba isimsiz dünya")
	}()
	//time.Sleep(time.Second * 5)
	fmt.Scanln() // --> bu scanln neyi yakalıyor ? **= main goroutune'in bitmesini önlüyor, time.sleep ile de yapılabilir.

}

func bulucuA(dizi []string) {
	for i := 0; i < len(dizi); i++ {
		if dizi[i] == "Flash" {
			fmt.Println("Bulucu A: " + dizi[i] + " buldu")
		}
		time.Sleep(time.Second)
	}
}

func bulucuB(dizi []string) {
	for i := 0; i < len(dizi); i++ {
		if dizi[i] == "Flash" {
			fmt.Println("Bulucu B: " + dizi[i] + " buldu")
		}
		time.Sleep(time.Second)
	}
}

func hello() {
	fmt.Println("hello fırladım")
}
