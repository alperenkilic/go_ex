package main

import (
	"fmt"
	"sync"
)

func main() {
	messages := make(chan string) // making channel
	var wg sync.WaitGroup         // making WaitGroup
	wg.Add(1)                     // adding markdown
	go func() {                   // running func with another goroutine
		messages <- "sended ping" // string send to messages channel
		defer wg.Done()           // lastly WaitGroup received "done"
	}() // last `()` for --> go func
	msg := <-messages // messages channel transfering "sended ping" string to msg
	fmt.Println(msg)  // şov işte
	wg.Wait()         // waiting to go routine
}
