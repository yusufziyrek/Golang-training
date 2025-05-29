package types

type TodoCreateDTO struct {
	Title string `json:"title" validate:"required"`
}

type TodoResponseDTO struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}