package main

import (
	"fmt"
	"time"
)

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result
}
func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
	}
	ch <- result
}

/*
A select statement blocks until at least one of its cases can proceed.
The default case is useful in preventing deadlocks -- without it the
select would wait for a channel forever, crashing the program if none
of the channels received data.
*/
func main() {
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)

	for { // waiting for one of the channels to receive data:
		select {
		case x := <-evenCh:
			fmt.Println(x)
			return
		case y := <-sqCh:
			fmt.Println(y)
			return
		default:
			fmt.Println("Nothing available")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
