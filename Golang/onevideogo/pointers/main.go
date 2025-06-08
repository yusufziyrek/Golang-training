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

	x := 10
	changeValue(&x)
	fmt.Println(x) // 15

}

func changeValue(value *int) {
	*value = *value + 5
}
