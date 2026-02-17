package main

import (
	"fmt"
	"net/http"
	"simple-product-api/config"
	"simple-product-api/handler"
	"simple-product-api/route"
	"simple-product-api/repository"
	"simple-product-api/service"
)

func main() {
	//init DB
	db := config.Connect()
	defer db.Close()

	//Dependency injection
	repository := repository.ProductRepo{DB: db}
	service := service.ProductService{Repo: &repository}
	handler := handler.ProductHandler{Service: &service}
	prodRoute := route.ProductRoute{Handler: &handler}

	mux := http.NewServeMux()
	prodRoute.Product(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
