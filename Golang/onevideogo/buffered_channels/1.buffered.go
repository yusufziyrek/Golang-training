package main

import (
	"fmt"
	"time"
)

/*
		Buffered channel kullanımında taşınacak veri miktarı belirtilir !!
		Buffered channel, veri giriş çıkışını sıkı sıkıya kontrol etmez !!
	   	Fazla veri girilirse veya olmayan veri alınmaya çalışılırsa deadlock olur !!
*/
func main() {

	bufferedChannel := make(chan int, 4)

	go func() {
		for i := 1; i < 10; i++ {
			bufferedChannel <- i
			fmt.Println("Sent data : ", i)
		}

		close(bufferedChannel)
	}()

	time.Sleep(3 * time.Second)

	for data := range bufferedChannel {
		fmt.Println("Received data : ", data)
		time.Sleep(1 * time.Second)
	}

}
