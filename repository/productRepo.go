package repository

import (
	"database/sql"
	"fmt"
	"simple-product-api/models"
)

type ProductRepo struct {
	DB *sql.DB
}

func (pr *ProductRepo) GetProduct() ([]models.Product, error) {
	//Alur : Generate query, return domain struct

	var data []models.Product

	rows, err := pr.DB.Query("Select * from product")
	if err != nil {
		return []models.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var rowData models.Product

		if err = rows.Scan(&rowData.Id, &rowData.Namaprod, &rowData.Kategori, &rowData.Price, &rowData.Stock); err != nil {
			return []models.Product{}, err
		}
		data = append(data, rowData)
	}
	//semua aman
	return data, nil
}

func (pr *ProductRepo) InsertProduct(req models.ProductRequest) (models.Product, error) {
	//Alur : Jalanin query, return domain struct (ngamnbil id dari hasil auto increment table)

	//Tampungan domain struct
	product := models.Product{
		Namaprod: req.Namaprod,
		Kategori: req.Kategori,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	//query row exec query dan return data including ID pake returning, sebagai return value
	query := "Insert into product (namaprod, kategori, price, stock) values (?,?,?,?)"

	//logikanya tu karna domain struct punya value sama aja, disini hasil queryrow return ID
	//karna cukup butuh mappingan last inserted id untuk generate id product baru
	result, err := pr.DB.Exec(query, req.Namaprod, req.Kategori, req.Price, req.Stock)
	if err != nil {
		return models.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}
	product.Id = int(id)

	//artinya aman
	return product, nil
}

func (pr *ProductRepo) UpdateProductByID(id int, req models.ProductRequest) (models.Product, error) {
	//Alur : Jalanin query, return domain struct (semua info property udh dari request)

	product := models.Product{
		Id:       id, //diambil dari url, dipassing sebagai param
		Namaprod: req.Namaprod,
		Kategori: req.Kategori,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	query := "update product set namaprod=?, kategori=?, price=?, stock=? where id = ?"
	res, err := pr.DB.Exec(query, product.Namaprod, product.Kategori, product.Price, product.Stock, product.Id)
	if err != nil {
		return models.Product{}, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return models.Product{}, err
	}
	if rows == 0 {
		return models.Product{}, err
	}

	//berarti aman
	return product, nil
}

func (pr *ProductRepo) DeleteProductByID(id int) (models.Product, error) {
	//Alur : Jalanin query, return domain struct (ngamnbil id dari hasil auto increment table)

	//Tampungan domain struct
	product := models.Product{
		Id: id,
	}

	//query select based id, untuk dpt info deleted baru jalanin delete query
	err := pr.DB.QueryRow("select namaprod,kategori,price,stock from product where id = ?", id).
		Scan(&product.Namaprod, &product.Kategori, &product.Price, &product.Stock)

	fmt.Println(product.Id, product.Namaprod, product.Kategori, product.Price, product.Stock)
	if err != nil {
		return models.Product{}, err
	}

	result, err := pr.DB.Exec("Delete from product where id = ?", id)
	if err != nil {
		return models.Product{}, err
	}
	rowsAff, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, err
	}
	if rowsAff == 0 {
		return models.Product{}, err
	}
	//artinya aman
	return product, nil
}
