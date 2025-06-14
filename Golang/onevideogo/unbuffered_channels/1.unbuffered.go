package main

// Main go routine herhangi bir channel bilgisini beklemez !!
/* func main() {

	myChannel := make(chan string)
	done := make(chan bool)

	go func() {
		message := <-myChannel // Kanalı dinler
		fmt.Println(message)
		done <- true
	}()

	go func() {
		myChannel <- "Hello World" // Kanala bilgi gönderir
	}()

	<-done // Bilginin alınmasını bekler ve akışı bloklar

	fmt.Println("End of the main")

}

*/
