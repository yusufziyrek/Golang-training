package main

import (
	"errors"
	"fmt"
)

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{text: "Ürün ismi boş olamaz!", code: "1001"}
	}
	if product.price < 0 {
		return ValidationError{text: "Ürün fiyatı sıfırdan büyük olmalı!", code: "1002"}
	}

	fmt.Println("Ürün eklendi")
	return nil

}

func main() {

	/* fileData, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	} */

	/* result, err := divide(3, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	} */

	productService := ProductService{}
	err := productService.Add(Product{id: 1, name: "asdasdasd", price: -10})
	if err != nil {
		fmt.Println(err)
	}

}

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata : %s , Kod: %s", validationError.text, validationError.code)
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("payda sıfır olamaz")
	}

	return x / y, nil
}
