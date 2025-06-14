package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/* Main işlem akışından ayrılan fonksiyonlar main işlemi devam ederken yan dallardan çalışmaya devam eder.
	   Go keyword kullanarak ilgili işlemi farklı bir yaşam döngüsüne çekeriz yani main akıştan ayırırız.
	*/

	/* go f1() // go keyword ile asenkron işlemi başlatırız
	go f2()
	fmt.Println("End of the main")
	time.Sleep(1 * time.Second) // time sleep kullanmak her zaman iyi bir fikir değil
	*/

	// Grup go routine beklemesi, sağlıklı asenkron işlemler sağlar
	/* wg := sync.WaitGroup{}
	wg.Add(2) // Kaç tane go routine olacağını belirtiriz

	go func() {
		defer wg.Done() // Asenkron işlemin bittiği bilgisini sayaca bildirir
		fmt.Println("f1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("f2")
	}()

	// !! wg.Add() ile belirttiğimizden fazla go routine işlemi varsa sistem deadlock olur !!
	wg.Wait() // Go routinlerin bitmesini bekler

	fmt.Println("End of the main") */

	syncFlow()
	asyncFlow()
}

func syncFlow() {

	startTime := time.Now()

	func() {
		fmt.Println("f1")
		time.Sleep(3 * time.Second)

	}()

	func() {
		fmt.Println("f2")
		time.Sleep(3 * time.Second)

	}()

	func() {
		fmt.Println("f3")
		time.Sleep(3 * time.Second)

	}()

	fmt.Println("Passed time for sync : ", time.Since(startTime))

}

func asyncFlow() {

	startTime := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("f1")
		time.Sleep(3 * time.Second)

	}()

	go func() {
		defer wg.Done()
		fmt.Println("f2")
		time.Sleep(3 * time.Second)

	}()

	go func() {
		defer wg.Done()
		fmt.Println("f3")
		time.Sleep(3 * time.Second)

	}()

	wg.Wait()

	fmt.Println("Passed time for async : ", time.Since(startTime))

}
