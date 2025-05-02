package dal

type Todo struct {
	ID        int // Büyük harfle ID yazınca gorm bunu algılayacak ve otomatik olarak bunu attıracak !!
	Title     string
	Completed bool `gorm:"default:false"`
}
