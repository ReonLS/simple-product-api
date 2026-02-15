package route

import "net/http"
import "simple-product-api/handler"

//ngebuat instans of product handler
type ProductRoute struct{
	Handler *handler.ProductHandler
}

func (pr *ProductRoute) ProductRoutingByID(rw http.ResponseWriter, r *http.Request) {
	//alur komunikasi : main -> route -> handler -> db
	//handler buat instans db, route buat instans

	//cek method dari request terus switch berdasarkan case
	switch r.Method {
	case "PUT":
		pr.Handler.UpdateProductByID(rw, r)

	case "DELETE":
		pr.Handler.DeleteProductByID(rw, r)
	}
}