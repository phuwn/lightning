package model

// Product data model
type Product struct {
	Base
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Photo string  `json:"photo"`
}
