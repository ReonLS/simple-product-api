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

	//Konek instans DB handler dengan database
	handler := h.ProductHandler{DB: db}
	route := r.ProductRoute{Handler: &handler}

	http.HandleFunc("/product", route.ProductRouting)      //GET & POST
	http.HandleFunc("/product/", route.ProductRoutingByID) //PUT & DELETE

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
