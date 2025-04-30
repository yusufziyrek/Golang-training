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

	//s1 := student.New("1", "Yusuf", 23, "CENG", []string{"Math", "Machine Learning", "Computer Organization"})
	//s2 := student.New("2", "Akif", 22, "CENG", []string{"Math", "Machine Learning"})

	//s1.ShowInfo()
	//s2.ShowInfo()
	//s2.Rename("İlkay")
	//fmt.Println("s2 new name : ", s2.Name)

	student := student.CreateWithScan()
	student.ShowInfo()

}
