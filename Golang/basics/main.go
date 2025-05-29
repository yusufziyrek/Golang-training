package main

import (
	"fmt"

	"go_training.go/student"
	"go_training.go/training"
)

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

	//student := student.CreateWithScan()
	//student.ShowInfo()

	//switchProcess()

	// !! Interface kullanımı !!

	var shape training.Shape
	c := training.Circle{Radius: 5}
	shape = c
	fmt.Println("Dairenin Alanı:", shape.Area()) // Alan hesaplanır

	r := training.Rectangle{Width: 4, Height: 6}
	shape = r
	fmt.Println("Dikdörtgenin Alanı:", shape.Area()) // Alan hesaplanır

}

func switchProcess() {

	studentList := []student.Student{}
	var secim int

main_lopp:
	for {

		fmt.Println("\n! Yapmak istediğiniz işlemi giriniz !")
		fmt.Println(
			"1- Öğrenci eklemek\n" +
				"2- Öğrencileri listele\n" +
				"3- Dosyaya kaydet\n" +
				"4- Çıkış yap",
		)

		fmt.Scan(&secim)
		switch secim {
		case 1:
			newStudent := student.CreateWithScan()
			studentList = append(studentList, newStudent)
		case 2:
			if len(studentList) == 0 {
				fmt.Println("\n!! Listenizde öğrenci bulunmamakta !!")
			}
			for _, student := range studentList {
				student.ShowInfo()
			}
		case 3:
			student.SaveToFile(studentList)
		case 4:
			fmt.Println("!! Çıkış yaptınız !!")
			break main_lopp
		default:
			fmt.Println("!! Geçersiz işlem !!")
		}

	}
}
