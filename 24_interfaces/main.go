package main

import "fmt"

type kare struct {
	a, b int
}

type çember struct {
	r int
}

func (k kare) çevre() int {
	return (k.a + k.b) * 2
}

func (ç çember) çevre() int {
	return 2 * 3 * ç.r
}

type şekil interface {
	çevre() int
}

func interfaceFunc(ş şekil) int {
	return ş.çevre()
}

func main() {
	k := kare{a: 3, b: 4}
	ç := çember{r: 5}
	fmt.Println(interfaceFunc(k))
	fmt.Println(interfaceFunc(ç))
}
