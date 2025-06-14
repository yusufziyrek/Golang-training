package main

import "fmt"

func main() {
	myChannel := make(chan int)

	go func() {
		data := <-myChannel
		fmt.Println("First go routine took data : ", data)
	}()

	go func() {
		data := <-myChannel
		fmt.Println("Second go routine took data : ", data)
	}()

	myChannel <- 10
	fmt.Println("End of the main")
}
