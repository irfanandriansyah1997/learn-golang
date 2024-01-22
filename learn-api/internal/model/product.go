package model

type Product struct {
	Description string  `json:"description" validate:"required,min=5"`
	ID          uint    `json:"id"`
	Name        string  `json:"name" validate:"required,min=5"`
	Price       float64 `json:"price" validate:"required"`
}
