package main

import "fmt"

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

	/* book := &Book{desi: 10}
	book2 := &Book{desi: 20}

	fmt.Println(book.ShippingCost())
	fmt.Println(book2.ShippingCost()) */

	electronic1 := &Electronic{desi: 20}
	fmt.Println(electronic1.ShippingCost())

	electronics := []Electronic{{desi: 10}, {desi: 20}}
	calculateTotalShippingCostOfElectronics(electronics)

	/* books := []Book{{desi: 10}, {desi: 20}}
	fmt.Println(calculateTotalShippingCostOfBooks(books)) */

}

func calculateTotalShippingCostOfBooks(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.ShippingCost()
	}

	return total
}

func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0
	for _, electronic := range electronics {
		total += electronic.ShippingCost()
	}

	return total
}
