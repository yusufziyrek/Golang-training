package main

import "fmt"

func main() {

	func() {
		fmt.Println("Embedded function")
	}() // Bulunduğu yerde çalışması için () konulur

	// Bu fonksiyonları değişkenlere atayabiliriz
	add := func(x int, y int) int {
		return x + y
	}

	multiply := func(x int, y int) int {
		return x * y
	}

	a, b := calculator(3, 5, add, multiply)
	fmt.Println(a, b)

}

// Parametre olarak başka bir fonksiyon da ekleyebiliriz
func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)

}
