package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	startTime := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("A işlemi yapılıyor..")
		time.Sleep(3 * time.Second)
		fmt.Println("A işlemi bitti !")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("B işlemi yapılıyor..")
		time.Sleep(1 * time.Second)
		fmt.Println("B işlemi bitti !")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("C işlemi yapılıyor..")
		time.Sleep(2 * time.Second)
		fmt.Println("C işlemi bitti !")
	}()

	wg.Wait()

	fmt.Println("Finish time : ", time.Since(startTime))

}
