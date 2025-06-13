package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

func (book *Book) ShippingCost() int {

	return 5 + book.desi*2

}

func (electronic *Electronic) ShippingCost() int {

	return 10 + electronic.desi*3

}

func main() {

	/* var product IShippable

	product = &Book{desi: 10}
	fmt.Println(product.ShippingCost())

	product = &Electronic{desi: 10}
	fmt.Println(product.ShippingCost()) */

	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Book{desi: 20},
		&Electronic{desi: 20},
	}

	fmt.Println(calculateTotalShippingCost(products))

	/* book := &Book{desi: 10}
	book2 := &Book{desi: 20}

	fmt.Println(book.ShippingCost())
	fmt.Println(book2.ShippingCost())

	electronic1 := &Electronic{desi: 20}
	fmt.Println(electronic1.ShippingCost())

	electronics := []Electronic{{desi: 10}, {desi: 20}}
	calculateTotalShippingCostOfElectronics(electronics)

	books := []Book{{desi: 10}, {desi: 20}}
	fmt.Println(calculateTotalShippingCostOfBooks(books)) */

}

func calculateTotalShippingCost(products []IShippable) int {
	total := 0

	for _, product := range products {
		total += product.ShippingCost()
	}

	return total

}
