package main

import (
	"fmt"
)

func main() {
	switch g := 12; g {
	case 1:
		fmt.Println("değer 1")
	case 5:
		fmt.Println("değer 5")
	case 6:
		fmt.Println("değer 6")
	case 7:
		fmt.Println("değer 7")
	case 8:
		fmt.Println("değer 8")
	case 12:
		fmt.Println("değer 12")
	}
}
