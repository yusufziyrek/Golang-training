package main

import "fmt"

func main() {

	var exam float32

	fmt.Print("Please enter your exam note : ")
	fmt.Scan(&exam)

	if 90 <= exam && exam <= 100 {
		fmt.Println("Your note AA !!")
	} else if 70 <= exam && exam <= 89 {
		fmt.Println("Your note BB !!")
	} else if 50 <= exam && exam <= 69 {
		fmt.Println("Your note CC !!")
	} else if 30 <= exam && exam <= 49 {
		fmt.Println("Your note DD !!")
	} else if 0 <= exam && exam <= 29 {
		fmt.Println("Your note FF !!")
	} else {
		fmt.Print("Input error !!")
	}

}
