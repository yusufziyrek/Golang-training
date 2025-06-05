package main

import "fmt"

func main() {

	// Önceden bilinen değerler ile başlangıç
	persons := map[string]int{
		"Yusuf":   1,
		"İlkay":   2,
		"Zeyna":   3,
		"Karamel": 4,
	}
	persons["Esma"] = 5

	fmt.Println(persons)

	// Boş map oluşturul sonradan doldurulur
	persons2 := make(map[string]int)
	persons2["Akif"] = 1
	persons2["Ilgın"] = 2

	fmt.Println(persons2)

}
