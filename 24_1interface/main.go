package main

import (
	"fmt"
	"strconv"
)

// struckt

type car struct {
	Brand string
	Model string
	Color string
	Speed float64
	Price float64
}

type specialProduction struct {
	Special bool
}

type ferrari struct {
	car
	specialProduction
}

type merc struct {
	car
	specialProduction
}

type carface interface {
	run() bool
	stop() bool
	information() string
}

func (_ ferrari) run() bool { // girdi almayacak girdi tipi belli.
	return true
}

func (_ ferrari) stop() bool {
	return true
}

func (x *ferrari) information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "color: " + x.Color + "\n" + "\t" + "speed: " + strconv.FormatFloat(x.Speed, 'g', -1, 64) + "\n" + "\t" + "price: " + strconv.FormatFloat(x.Price, 'g', -1, 64) + "millons"
	add := "yes"
	if x.Special {
		ret += "\n" + "\t" + "special : " + add
	}
	return ret
}

func (_ merc) run() bool { // girdi almayacak girdi tipi belli.
	return true
}

func (_ merc) stop() bool {
	return true
}

func (x *merc) information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "color: " + x.Color + "\n" + "\t" + "speed: " + strconv.FormatFloat(x.Speed, 'g', -1, 64) + "\n" + "\t" + "price: " + strconv.FormatFloat(x.Price, 'g', -1, 64) + "millons"
	add := "yes"
	if x.Special {
		ret += "\n" + "\t" + "special : " + add
		return ret
	}
	return ret
}
func main() {
	f1 := new(ferrari)
	f1.Brand = "ferrari"
	f1.Color = "kırmızı"
	f1.Model = "c200"
	f1.Speed = 250
	f1.Price = 0.4
	f1.Special = false
	fmt.Println(f1.information())

	m1 := new(merc)
	m1.Brand = "mercedes"
	m1.Special = true
	m1.Color = "beyaz"
	m1.Price = 0.5
	m1.Speed = 320
	m1.Model = "f30"
	fmt.Println(m1.information())

	fmt.Printf("\n \n \n")

	fmt.Println(f1.stop())
	fmt.Println(m1.run())
}
