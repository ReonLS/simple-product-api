package route

import "net/http"
import "simple-product-api/handler"

//ngebuat instans of product handler
type Route struct{
	ProdHandler *handler.ProductHandler
	UserHandler *handler.UserHandler
}

//centralized handler func for /product
func (pr *Route) Product (mux *http.ServeMux){
	mux.HandleFunc("GET /product", pr.ProdHandler.GetProduct)
	mux.HandleFunc("POST /product", pr.ProdHandler.InsertProduct)
	mux.HandleFunc("PUT /product/{id}", pr.ProdHandler.UpdateProductByID)
	mux.HandleFunc("DELETE /product/{id}", pr.ProdHandler.DeleteProductByID)
}

func (uh *Route) User (mux *http.ServeMux){
	mux.HandleFunc("GET /user", uh.UserHandler.GetAllUsers)
}
