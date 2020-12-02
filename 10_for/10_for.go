package main

import (
	"fmt"
	"time"
)

func main() {

	for x := 0; true; x++ {

		if x%3 == 0 {
			continue // continue döngünün başına gider-- break sonuna
		}
		time.Sleep(1 * time.Second)
		fmt.Println(x)
	}
}
