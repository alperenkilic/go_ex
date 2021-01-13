package main

import "fmt"

type dikdortgen struct {
	uzunKenar float64
	kısaKenar float64
}

func (s dikdortgen) cevre() float64 {
	return 2 * (s.kısaKenar + s.uzunKenar)
}

func (s dikdortgen) alan() float64 {
	return s.kısaKenar * s.uzunKenar
}

type daire struct {
	yariCap float64
}

func (s daire) cevre() float64 {
	return 2 * s.yariCap * 3.14
}

func (s daire) alan() float64 {
	return s.yariCap * s.yariCap * 3.14
}

type shape interface {
	cevre() float64
	alan() float64
}

func interfaceFunc(s shape) {

	fmt.Println(s.alan())
	fmt.Println(s.cevre())
}

func main() {
	d1 := dikdortgen{uzunKenar: 3.0, kısaKenar: 4.0}

	fmt.Println(d1.cevre())
	fmt.Println(d1.alan())

	c1 := daire{yariCap: 4}
	fmt.Println(c1.cevre())
	fmt.Println(c1.alan())
	fmt.Println()
	fmt.Println()

	interfaceFunc(c1)
	interfaceFunc(d1)

}
