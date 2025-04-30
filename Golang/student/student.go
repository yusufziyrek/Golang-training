package student

import (
	"fmt"
	"strings"
)

type Student struct {
	Id         string
	Name       string
	Age        int
	Department string
	Lectures   []string
}

// Constructor metod
func New(id string, name string, age int, department string, lectures []string) Student {

	return Student{
		Id:         id,
		Name:       name,
		Age:        age,
		Department: department,
		Lectures:   lectures,
	}

}

// Receiver kullanımı - Oluşturulan obje üzerinden işlem yapabilme olanağı sağlar !!
// Standart receiver obje kopyası üzerinde işlem yapar. Objenin özelliklerini değiştiremez !!
func (s Student) ShowInfo() {
	fmt.Println(
		"-------------\n",
		"Id : ", s.Id, "\n",
		"Name : ", s.Name, "\n",
		"Age : ", s.Age, "\n",
		"Department : ", s.Department, "\n",
		"Lectures : ", s.Lectures,
	)
}

// Orjinal obje üzerinde değişiklik yapmak için pointer receiver kullanımı gerekli !! -- (*Student) --
func (s *Student) Rename(newName string) {
	s.Name = newName
}

func CreateWithScan() Student {

	var id, name, department, lecturesText string
	var age int
	fmt.Print("Your id : ")
	fmt.Scan(&id)
	fmt.Print("Your name : ")
	fmt.Scan(&name)
	fmt.Print("Your age : ")
	fmt.Scan(&age)
	fmt.Print("Your department : ")
	fmt.Scan(&department)
	fmt.Print("Your lectures with , : ")
	fmt.Scan(&lecturesText)                      
	lectures := strings.Split(lecturesText, ",") // Virgülle ayrılmış dersler

	newStudent := Student{
		Id:         id,
		Name:       name,
		Age:        age,
		Department: department,
		Lectures:   lectures,
	}

	return newStudent

}
