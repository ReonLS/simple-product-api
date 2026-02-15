package main

import "fmt"
import "simple-product-api/handler"
import "simple-product-api/config"
import "net/http"

func main(){
	//init DB
	db := config.Connect()
	defer db.Close()

	//Konek instans DB handler dengan database
	h := handler.ProductHandler{DB: db}
	http.HandleFunc("/product", h.GetProduct)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}