package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"required"`
}
