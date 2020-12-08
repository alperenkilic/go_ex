package main

import "fmt"

// sayı dizisinden tek sayıları ve çift sayıları ayırma
func main() {

	numbers := []int{2, 3, 4, 12, 11, 2, 3, 5, 12, 2}
	//evenNumbers := make([]int, 10)  ---> make ile yapınca bir şey değişmedi
	evenNumbers, oddNumbers := []int{}, []int{} //--> boş slice'lar oluşturdum
	//var counter int = 0
	for _, v := range numbers { //-->numbers slice'ının içide dolandım
		if v%2 == 0 {
			//evenNumbers[counter] = v  ---> counter ile eklemeye çalıştım bu bir hataydı.
			evenNumbers = append(evenNumbers, v) //--> bütün olay APPEND ile eklemek. Append direkt sona ekliyor.
			//counter = counter + 1
		} else {
			oddNumbers = append(oddNumbers, v)
		}
	}
	fmt.Print("Even Numbers -->")
	fmt.Println(evenNumbers)
	fmt.Print("Odd Numbers -->")
	fmt.Println(oddNumbers)
}
