package models

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryID   *int   `json:"category_id"`
	CategoryName string `json:"category_name"`
}
