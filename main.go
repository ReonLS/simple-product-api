package main

import (
	"fmt"
	"net/http"
	c "simple-product-api/config"
	h "simple-product-api/handler"
	r "simple-product-api/route"
)

func main() {
	//init DB
	db := c.Connect()
	defer db.Close()

	//Dependency injection
	prodHand := h.ProductHandler{DB: db}
	prodRoute := r.ProductRoute{Handler: &prodHand}

	mux := http.NewServeMux()
	prodRoute.Product(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
