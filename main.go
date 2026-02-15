package main

import "fmt"
import h "simple-product-api/handler"
import c "simple-product-api/config"
import r "simple-product-api/route"
import "net/http"

func main(){
	//init DB
	db := c.Connect()
	defer db.Close()

	//Konek instans DB handler dengan database
	handler := h.ProductHandler{DB: db}
	route := r.ProductRoute{Handler: &handler}

	http.HandleFunc("/product", handler.GetProduct) //GET
	http.HandleFunc("/", handler.InsertProduct) //POST
	http.HandleFunc("/product/", route.ProductRoutingByID) //PUT & DELETE

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}