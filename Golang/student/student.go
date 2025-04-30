package student

import (
	"fmt"
	"strings"
)

var nextId int

type Student struct {
	Id         int
	Name       string
	Age        int
	Department string
	Lectures   []string
}

// Constructor metod
func New(name string, age int, department string, lectures []string) Student {

	nextId++

	return Student{

		Id:         nextId,
		Name:       name,
		Age:        age,
		Department: department,
		Lectures:   lectures,
	}

}

// Receiver kullanımı - Oluşturulan obje üzerinden işlem yapabilme olanağı sağlar !!

// Value receiver, obje kopyası üzerinde işlem yapar. Objenin özelliklerini değiştiremez !!
func (s Student) ShowInfo() {
	fmt.Printf(
		"-------------\n"+
			"Id : %d\n"+
			"Name : %s\n"+
			"Age : %d\n"+
			"Department : %s\n"+
			"Lectures : %v\n"+
			"-------------\n",
		s.Id, s.Name, s.Age, s.Department, s.Lectures,
	)
}

// Pointer receiver, orjinal obje üzerinde değişiklik yapabilme imkanı sağlar !! -- (*Student) --
func (s *Student) Rename(newName string) {
	s.Name = newName
}

func CreateWithScan() Student {

	var name, department, lecturesText string
	var age int

	fmt.Print("Your name : ")
	fmt.Scan(&name)
	fmt.Print("Your age : ")
	fmt.Scan(&age)
	fmt.Print("Your department : ")
	fmt.Scan(&department)
	fmt.Print("Your lectures with , : ")
	fmt.Scan(&lecturesText)
	lectures := strings.Split(lecturesText, ",")

	nextId++

	newStudent := Student{
		Id:         nextId,
		Name:       name,
		Age:        age,
		Department: department,
		Lectures:   lectures,
	}

	return newStudent

}
