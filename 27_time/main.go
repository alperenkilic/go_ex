package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Weekday())

	diff := now.Sub(then) // şimdi ve diğer tarihin farkı
	p(diff)

	p(diff.Hours())

	p(then.Add(diff))
	p(then.Add(-diff))
}
