package dal

import (
	"gorm.io/gorm"
	"rest_api.go/database"
)

type Todo struct {
	ID        int // Büyük harfle ID yazınca gorm bunu algılayacak ve otomatik olarak bunu attıracak !!
	Title     string
	Completed bool `gorm:"default:false"`
}

func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(&todo)
}
