package models

type Product struct{
	Id int
	Namaprod string
	Kategori string
	Price float64
	Stock int
}

type ProductRequest struct{
	Namaprod string `json:"namaprod" binding:"required"`
	Kategori string `json:"kategori" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Stock int `json:"stock" binding:"required,gt=0"`
}

type ProductResponse struct{
	Id int `json:"id"`
	Namaprod string `json:"namaprod"`
	Kategori string `json:"kategori"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}

