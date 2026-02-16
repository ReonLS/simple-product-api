package route

import "net/http"
import "simple-product-api/handler"

//ngebuat instans of product handler
type ProductRoute struct{
	Handler *handler.ProductHandler
}

//centralized handler func for /product
func (pr *ProductRoute) Product (mux *http.ServeMux){
	mux.HandleFunc("GET /product", pr.Handler.GetProduct)
	mux.HandleFunc("POST /product", pr.Handler.InsertProduct)
	mux.HandleFunc("PUT /product/{id}", pr.Handler.UpdateProductByID)
	mux.HandleFunc("DELETE /product/{id}", pr.Handler.DeleteProductByID)
}
