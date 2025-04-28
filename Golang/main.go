package main

import (
	"go_training.go/student"
)

//"go_training.go/training"

func main() {

	// Paket kullanımları

	/*training.Degiskenler()
	training.PrintIslemleri()
	training.Kosullar()
	training.Diziler()
	training.Mapler()
	training.Donguler()
	fmt.Println(training.AlanHesapla(3))
	*/

	// -- Pointer kullanımı --
	//training.Pointer()

	s1 := student.New("1", "Yusuf",23, "CENG", []string{"Math", "Machine Learning", "Computer Organization"})
	s2 := student.New("2", "Akif", 22, "CENG", []string{"Math", "Machine Learning"})

	student.ShowInfo(s1)
	student.ShowInfo(s2)
}
