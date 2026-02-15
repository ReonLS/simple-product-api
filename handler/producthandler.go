package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	m "simple-product-api/models"
)

// buat instans DB, jdi layer handler bisa exec query
type ProductHandler struct {
	DB *sql.DB
}

var baseurl = "http://localhost:8080"

// Fungsi utama responds to a request (getting all product)
func (ph *ProductHandler) GetProduct(rw http.ResponseWriter, r *http.Request) {

	//Alur: setheader -> DB Query -> Append Data -> convert to Json -> Response Writer
	rw.Header().Set("Content-Type", "application/json")
	data := []m.Product{}

	//ini rows contain instans data hasil query
	rows, err := ph.DB.Query("Select * from product")
	if err != nil {
		//critical error, generate response code server fault
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if rows.Err() != nil {
		//critical error, generate response code server fault
		http.Error(rw, rows.Err().Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		//tampungan per row, buat append ke data
		temp := m.Product{}

		if err := rows.Scan(&temp.Id, &temp.Namaprod, &temp.Kategori, &temp.Price, &temp.Stock); err != nil {
			//critical error, generate response code server fault
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		data = append(data, temp)
	}
	//kalo disini berarti proses aman, set status code success
	rw.WriteHeader(http.StatusOK)

	//Best Approach, more memory efficient
	err = json.NewEncoder(rw).Encode(&data)
	if err != nil {
		//server-side error
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	//previous approach
	// jsonpayload, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("Error: ", err.Error())
	// }
	// rw.Write(jsonpayload)
}
