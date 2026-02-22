package repository

import (
	"database/sql"
	"simple-product-api/models"
)

type UserRepo struct {
	DB *sql.DB
}

func (ur *UserRepo) GetAllUsers()([]models.User, error){
	//Alur: Execute query, return domain struct
	var data []models.User

	rows, err := ur.DB.Query("Select * from user")
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var row models.User

		if err := rows.Scan(&row.Id, &row.Name, &row.Email, &row.IsAdmin); err != nil{
			return []models.User{}, err
		}
		data = append(data, row)
	}
	//udh aman
	return data, nil
}

func (ur *UserRepo) CreateUsers(model models.UserRequest)([]models.User, error) {
	ur.DB.Exec("insert into user ")
}