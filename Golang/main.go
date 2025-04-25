package main

import (
	"fmt"
)

func main() {

	//degiskenler()
	//printIslemleri()
	//kosullar()
	//diziler()
	donguler()
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

func kosullar() {

	var not int = 13

	if 90 <= not && not <= 100 {
		fmt.Println("AA")
	} else if 70 <= not && not <= 89 {
		fmt.Println("BB")
	} else if 50 <= not && not <= 69 {
		fmt.Println("CC")
	} else if 30 <= not && not <= 49 {
		fmt.Println("DD")
	} else {
		fmt.Println("FF")
	}

}

func diziler() {

	// Array "Statik yapı"
	isimler := [3]string{"Yusuf", "Akif", "Cengiz"}
	fmt.Println("İsimler : ", isimler)
	fmt.Println("İlk eleman : ", isimler[0])
	isimler[2] = "Mert"
	fmt.Println(isimler)

	fmt.Println("------")

	// Slice "Dinamik yapı"
	maaslar := []float32{10.9, 15.2, 35.2}
	fmt.Println("Maaşlar : ", maaslar)
	maaslar = append(maaslar, 69.3) // Veri ekleme
	fmt.Println(maaslar)
	maaslar = maaslar[0:3] // 4.elemana kadar al gerisini sil
	fmt.Println("Yeni maaşlar", maaslar)
	fmt.Println(len(maaslar)) // Uzunluk

}

func donguler() {

	// Standart for
	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}

	// While benzeri for
	count := 5
	for 0 < count {
		fmt.Println(count)
		count--
	}

	// Foreach benzeri for
	array := []string{"Yusuf", "Akif", "Doğukan", "Cengiz"}
	for index, value := range array {
		fmt.Println("Index : ", index, "Value : ", value)

	}
}
