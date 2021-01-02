package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("adınızı giriniz: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("HATA ! =")
		fmt.Println(err)
	} else {
		fmt.Println("hoşgeldin " + str)
	}

	fmt.Print("yaşını gir: ")
	str, _ = reader.ReadString('\n')
	yas, _ := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(yas)
	}

}
