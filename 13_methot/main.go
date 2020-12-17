package main

import (
	"fmt"
	"os"
)

/* func main() {
	var moment time.Time = time.Now()
	fmt.Println(moment)
} */

/* func main() {
	fmt.Print("değer giriniz ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println(value)
}
*/

/* func main() {

	sonuc, err := sRoot(-77)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sonuc)
	}
}

func sRoot(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("hata --> nega değer girdin")
	}
	return math.Sqrt(x), nil
} */

func main() {

	file, err := os.Open("test.go")
	if err != nil {
		fmt.Println(file, "dosya bulunamadı")
	} else {
		fmt.Println("dosyayı açtık", file)
	}
}
