package main

import (
	"fmt"
	"sync"
)

// goroutinelerin birbirini beklemesi.

var wg sync.WaitGroup

func main() {
	wg.Add(1)   // goroutine ekliyoruz
	go printX() //--> içinde Done var, fonkisoyn işini bitiriyor.
	wg.Wait()   // önceki go routinein işini bitirmesini bekliyor.
	wg.Add(1)
	go printY()
	wg.Wait()
}

func printX() {
	for i := 0; i < 10; i++ {
		fmt.Print("x")
	}
	wg.Done()
}

func printY() {
	for i := 0; i < 10; i++ {
		fmt.Printf("y")
	}
	wg.Done()
}
