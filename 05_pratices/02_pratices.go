package main

import (
	"fmt"
	"math"
)

/* func main() {

	var puan = 50
	var ates float64 = 36.6								typeConversion'da küçük olanı büyüğe çevir
	fmt.Println(ates + float64(puan))
} */

/* func main() {
	var x, y int = 3, 4
	fmt.Println(x, y)

	x, y = y, x												değer değiştirme x,y = y,x
	fmt.Println(x, y)
} */

/* func main() {
	var sayi int = 65

	x := strconv.Itoa(sayi)
	fmt.Printf("%v  %T\n", sayi, sayi)
	fmt.Printf("%v  %T\n", x, x)
} */

/* func main() {
	x := 1
	fmt.Println(x)
	x++
	fmt.Println(x)
	fmt.Println(x--)  //--> 1 satırda 2 statement olaamz.
	// statemen emit, expression ifade.
} */

func main() {
	x := math.Pow(2, 3) // constlar compile time'da tanımlanır, variablelar ise run timeda
	fmt.Println(x)
	//// mat.pow kütüphanesi runtime da çalışacağı için compile time'da x tanımlanamaz, pow hesaplanmadan x e değer veirlmiyor
}
