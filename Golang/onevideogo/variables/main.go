package main

import (
	"fmt"
)

func main() {

	/*var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Iphone"
	quantity = 10
	discount = 0.37
	isInStock = true

	// reflect.TypeOf() değişken türünü belirtir
	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))
	*/

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Iphone"
	quantity = 10
	discount = 0.37
	isInStock = true

	fmt.Printf("Product name : %s , Quantity : %d , Discount : %f , IsInStock : %t", productName, quantity, discount, isInStock)

}
