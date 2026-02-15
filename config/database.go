package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	//open connection ke database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/katalog")

	if err != nil {
		fmt.Println("Database Gagal Connect: ", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("DB enggak respon: ", err)
	}

	fmt.Println("Berhasil Konek")
	return db
}
