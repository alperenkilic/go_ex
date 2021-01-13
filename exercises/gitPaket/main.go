package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	fmt.Println("renksiz renk neden renktir ?")
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")
	color.Green("golang komutu")
	color.Blue("Prints %s in blue.", "text")

}
