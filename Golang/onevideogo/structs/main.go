package main

import "fmt"

func main() {

	// RECEİVER, FONKSİYONU STRUCT İLE İLİŞKİLENDİRİR

	var customer1 = Customer{id: 1, name: "Yusuf", age: 23,
		adress: Adress{city: "İstanbul", district: "Avcılar"}}
	// var customer2 = Customer{id: 2, name: "İlkay", age: 21}

	/* customer1.age = 24
	fmt.Println(customer1)
	fmt.Println(customer2) */

	customer1.ToString()
	customer1.changeName("Yeni isim")
	customer1.ToString()

}

// Pointer receiver (orijinal değeri değiştirir)
func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

type Customer struct {
	id     int64
	name   string
	age    int
	adress Adress
}

type Adress struct {
	city     string
	district string
}

// Value Receiver kullanımı (orjinal değer değiştirilmez)
func (customer Customer) ToString() {
	fmt.Printf("Name : %s , Age : %d\n", customer.name, customer.age)
}
