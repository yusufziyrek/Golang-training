package main

import "fmt"

func main() {

	// İşaretlenen işlemi fonksiyon bitimine erteler
	// Stack yapısı vardır. İlk giren son çıkar
	/* defer fmt.Println("1. Defer")
	defer fmt.Println("2. Defer")
	defer fmt.Println("3. Defer")

	fmt.Println("Main Fonksiyonu") */

	/* var condition = true
	defer fmt.Println("1.Defer")
	if condition {
		return
	}

	defer fmt.Println("2.Defer") */

	/* x := 10
	y := 20

	// Defer kullanıldığı anda ki değerleri korur
	defer fmt.Println("x:", x)
	defer fmt.Println("y:", y)

	x = 100
	y = 200

	fmt.Println("x:",x)
	fmt.Println("y:",y) */

	var condition = true

	defer cleanup()

	if condition {
		defer panic("An error occurred")
	}

}

func cleanup() {
	fmt.Println("Cleanup worked..")
}
