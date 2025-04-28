package student

import "fmt"

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
func (s *Student) ShowInfo() {
	fmt.Println(
		"-------------\n",
		"Id : ", s.Id, "\n",
		"Name : ", s.Name, "\n",
		"Age : ", s.Age, "\n",
		"Department : ", s.Department, "\n",
		"Lectures : ", s.Lectures,
	)
}
