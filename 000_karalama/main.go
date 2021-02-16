package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedKanal := make(chan string, 3)
	go func() {
		bufferedKanal <- "birinci"
		fmt.Println("1. Gönderildi")
		bufferedKanal <- "ikinci"
		fmt.Println("2. Gönderildi")
		bufferedKanal <- "üçüncü"
		fmt.Println("3. Gönderildi")
	}()
	<-time.After(time.Second * 1)
	go func() {
		ilkOkunan := <-bufferedKanal
		fmt.Println("Alınıyor...")
		fmt.Println(ilkOkunan)
		ikinciOkunan := <-bufferedKanal
		fmt.Println(ikinciOkunan)
		ucuncuOkunan := <-bufferedKanal
		fmt.Println(ucuncuOkunan)
	}()
	<-time.After(time.Second * 5)
}
