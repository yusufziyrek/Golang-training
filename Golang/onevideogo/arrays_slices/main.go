package main

import "fmt"

func main() {

	/* var names [3]string

	names[0] = "Yusuf"
	names[1] = "İlkay"
	names[2] = "Karamel"

	for index, name := range names {
		fmt.Println("Name : ", name, "index : ", index)
	}

	*/

	/* var names = [3]string{"Yusuf","İlkay","Karamel"}
	fmt.Println(names)

	names[0] = "Yusufff"
	fmt.Println(names)

	fmt.Println(names[0:2])
	*/

	// Slice
	var names = []string{"Yusuf", "İlkay", "Zeyna", "Karamel"}
	fmt.Println(names)
	names = append(names, "Gülhal")
	fmt.Println(names)

}
