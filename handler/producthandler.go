package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	m "simple-product-api/models"
	"strconv"
	"strings"
)

// buat instans DB, jdi layer handler bisa exec query
// sebenarnya ini layer repo, yg diakses service, yg diakses handler tp for now okela
type ProductHandler struct {
	DB *sql.DB
}

// Fungsi utama responds to a request (getting all product)
func (ph *ProductHandler) GetProduct(rw http.ResponseWriter, r *http.Request) {

	//Alur: setheader -> DB Query -> Append Data -> convert to Json -> Response Writer
	rw.Header().Set("Content-Type", "application/json")
	data := []m.Product{}
	fmt.Println("Masuk GET")

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

func (ph *ProductHandler) InsertProduct(rw http.ResponseWriter, r *http.Request) {
	//Alur real life : nerima json -> decode dan simpan di tampungan, exec query, generate respon
	rw.Header().Set("Content-Type", "application/json")
	fmt.Println("Masuk POST")

	//membuat tampungan
	var product m.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()

	if err != nil {
		//dianggap client salah kirim input
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	//inserting product with products.property
	query := "Insert into product (namaprod, kategori, price, stock) values (?,?,?,?)"
	//kalo exec, jgn pake pointer tp pake value
	result, err := ph.DB.Exec(query, product.Namaprod, product.Kategori, product.Price, product.Stock)

	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	//return status http object berhasil dibentuk
	rw.WriteHeader(http.StatusCreated)

	//ngambil index id auto increment, diassign ke product untuk diencode
	index, err := result.LastInsertId()
	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
	}
	product.Id = int(index)

	//tampilin di endpoint sbg response request client
	json.NewEncoder(rw).Encode(&product)
}

func (ph *ProductHandler) UpdateProductByID(rw http.ResponseWriter, r *http.Request) {
	//alur : set header, take id from url.path, decode req.body, exec with decoded property, generate respons
	rw.Header().Set("Content-Type", "application/json")
	
	path := r.URL.Path //{/{id}}
	stringId := strings.TrimPrefix(path, "/product/") //{id}
	fmt.Println("Masuk PUT", stringId)

	id, err := strconv.Atoi(stringId)
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	//tampungan decode
	var update m.Product

	err = json.NewDecoder(r.Body).Decode(&update)
	defer r.Body.Close()

	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return
	}
	//update id tampungan, jd incase client kirim id di body, ketimpa ama id URL
	update.Id = id

	//db exec
	query := "update product set namaprod=?, kategori=?, price=?, stock=? where id = ?"
	res, err := ph.DB.Exec(query, update.Namaprod, update.Kategori, update.Price, update.Stock, id)
	if err != nil {
		//karna inputan user yg salah
		http.Error(rw, "", http.StatusBadRequest)
		return
	}
	rowsAff, err := res.RowsAffected()
	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}
	//berarti gk ada proses update
	if rowsAff == 0 {
		http.Error(rw, "", http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)

	//encode update untuk write ke stream
	err = json.NewEncoder(rw).Encode(&update)
	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}
}

func (ph *ProductHandler) DeleteProductByID(rw http.ResponseWriter, r *http.Request) {
	//alur : set header -> ambil ID dari url, decode, jalankan query, encode, response
	rw.Header().Set("Content-Type", "application/json")

	//generate id from path
	path := r.URL.Path
	idstring := strings.TrimPrefix(path, "/product/")
	id, err := strconv.Atoi(idstring)
	fmt.Println("Masuk DELETE", id)

	if err != nil{
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	//jalankan query
	res, err := ph.DB.Exec("delete from product where id = ?", id)
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}
	if rowsAff == 0 {
		http.Error(rw, "", http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
