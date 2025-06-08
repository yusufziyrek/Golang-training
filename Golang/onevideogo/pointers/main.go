package main

import "fmt"

func main() {

	/* var a int
	var p *int

	a = 10
	fmt.Println(a)

	p = &a
	*p = 20

	fmt.Println(a) */

	/* var a int = 10
	var b int
	var p *int

	b = a
	p = &a

	*p = 20
	fmt.Println(a, b) */

	/* x := 10
	changeValue(&x)
	fmt.Println(x) // 15 */

	var numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	changeValueOfArray(numbers)
	fmt.Println(numbers)

}

func changeValue(value *int) {
	*value = *value + 5
}

// Diziler ve slicelar referans tip olduğu için pointer kullanımı gerekmez
func changeValueOfArray(numbers []int) {
	numbers[0] = 100
}
