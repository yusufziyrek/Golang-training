package main

import "fmt"

func main() {

	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}*/

	/*arr := []string{
		"Yusuf",
		"İlkay",
		"Zeyna",
		"Karamel",
	}

	for index, value := range arr {
		fmt.Println(index, value)
	}*/

	text := "Golang"

	for _, character := range text {
		fmt.Printf("Character %c\n", character)
	}
}
