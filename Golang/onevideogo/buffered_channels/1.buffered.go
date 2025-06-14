package main

import "fmt"

/*
	Buffered channel kullanımında taşınacak veri miktarı belirtilir !!
	Buffered channel, veri giriş çıkışını sıkı sıkıya kontrol etmez !!
   	Fazla veri girilirse veya olmayan veri alınmaya çalışılırsa deadlock olur !!
*/
func main() {
	myChannel := make(chan int, 2)

	myChannel <- 1
	myChannel <- 2

	fmt.Println(<-myChannel)

}
