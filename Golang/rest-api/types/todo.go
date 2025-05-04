package types

type TodoCreateDTO struct {
	Title string `validate:"required"`
}
