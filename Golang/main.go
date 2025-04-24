package main

import "fmt"

func main() {

	//degiskenler()
	printIslemleri()
}

func degiskenler() {

	// String

	var name string = "Yusuf"
	var surname string = "Ziyrek"
	var school string

	school = "Biruni"

	fmt.Println("Name :" + name)
	fmt.Println("Surname : " + surname)
	fmt.Println("School : " + school)

	fmt.Println("----------")

	// Int

	var number int = 10
	var x = 5
	var a, b, c, d int = 1, 2, 3, 4

	fmt.Println(number)
	fmt.Println(x)
	fmt.Println(a, b, c, d)

	// Float

	var f1 float32 = 3.5
	var f2 float32 = 4.9

	fmt.Println(f1 + f2)

	// int ve float kullanarak islem yapmak icin donusum yapmaliyiz
	fmt.Println(f1 + float32(number))

	// Boolean
	//var isActive bool = false

	kelime := "Merhaba"
	sayi := 1

	fmt.Println(kelime)
	fmt.Println(sayi)

}

func printIslemleri() {

	// Standart print islemleri

	fmt.Print("Tek satır cümle \n")
	fmt.Println("Yeni satıra geçiren cümle")
	fmt.Println("Yeni satır cümlesi")

	// Formatlı cümle
	var name string = "Yusuf"
	var age int = 22
	fmt.Printf("Your name : %v , Your age : %v \n", name, age)

	// String degiskenlerini tırnak içine alır
	fmt.Printf("Your name : %q , Your age : %q \n", name, age)

	// Degiskenin tipini belirtir
	fmt.Printf("The type of the variable \"name\" is %T \n", name)

	// Ondalıklı sayıları bastırmaya yarar ve istediğimiz kadar hane alabiliriz
	var price float32 = 31.69
	fmt.Printf("The price is : %0.01f \n", price)

	// Formatlanmış metni değişkene atar "Sprintf"
	var text string = fmt.Sprintf("The price is : %0.03f", price)
	fmt.Println(text)

}
