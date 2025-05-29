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

func GetAllTodo(dest any) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest)

}

func GetTodoById(dest any, id any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ?", id).First(dest)
}
