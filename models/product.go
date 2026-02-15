package models

type Product struct{
	Id int `json:"id"`
	Namaprod string `json:"namaprod" binding:"required"`
	Kategori string `json:"kategori" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Stock int `json:"stock" binding:"required,gt=0"`
}