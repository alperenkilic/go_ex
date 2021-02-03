package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("anlık zaman %v \n", time.Now())
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now().Unix())

	fmt.Println("--------------")

	t := time.Date(2020, time.September, 9, 1, 2, 2, 0, time.UTC)
	fmt.Printf("tarih: %s\n", t)

	fmt.Println("--------------")

	now := time.Now()
	fmt.Printf("şu an: %s\n", now)
	fmt.Println("--------------")

	fmt.Println(now.Month(), "ayı")
	fmt.Println(now.Day(), ". gün ")
	fmt.Println("günlerden", now.Weekday())
	fmt.Println("--------------")

	yarın := now.AddDate(0, 0, -11)
	fmt.Println(yarın)
	fmt.Println("--------------")

	kısaHal := "1/5/2020"
	tarih := t.Format(kısaHal)
	fmt.Println(tarih)
	fmt.Println("--------------")

	fark := now.Sub(now)
	fmt.Println(fark)

}
