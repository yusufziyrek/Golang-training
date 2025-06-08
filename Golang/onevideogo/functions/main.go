package main

import (
	"fmt"
)

func main() {

	/* total := add(10, 30)
	fmt.Println(total) */

	/* total, difference, multiply  := calculation(10, 20)
	fmt.Println(total, difference, multiply) */

	/* var numbers = []int{10, 20, 2, 3, 5}
	fmt.Println(sum(numbers)) */

	fmt.Println(sumOfDynamicParameter(1, 2, 3))
	fmt.Println(sumOfDynamicParameter(1, 2, 3, 4, 5, 6))

}

func add(x int, y int) int {

	//fmt.Println(x+y)

	return x + y

}

// Fonksiyonlar geriye birden çok değer döndürebilir
func calculation(x int, y int) (int, int, int) {

	return x + y, x - y, x * y

}

func sum(numbers []int) int {

	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum

}

// İstediğimiz kadar parametre gönderebiliriz
func sumOfDynamicParameter(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum

}
