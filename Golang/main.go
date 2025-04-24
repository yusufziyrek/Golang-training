package main

import "fmt"

func main() {

	// degiskenler()

	prints()
}

func degiskenler(){

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

func prints(){

	// Standart print islemleri

	fmt.Print("Tek satır cümle \n")
	fmt.Println("Yeni satıra geçiren cümle")
	fmt.Println("Yeni satır cümlesi")
	
	// Formatlı cümle

	var name string = "Yusuf"
	var age int = 22

	fmt.Printf("Your name : %v , Your age : %v" ,name,age)
}

