package service

import (
	"simple-product-api/repository"
	"simple-product-api/models"
)

type UserService struct{
	Repo *repository.UserRepo
}

func ToAdminUserResponse(user models.User) (models.AdminUserResponse){
	return models.AdminUserResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		IsAdmin: user.IsAdmin,
	}
}

func ToUserResponse(user models.User) (models.UserResponse){
	return models.UserResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
	}
}

func (us *UserService) GetAllUsers()([]models.AdminUserResponse, error){
	var response []models.AdminUserResponse
	
	data, err := us.Repo.GetAllUsers()
	if err != nil {
		return []models.AdminUserResponse{}, err
	}

	for _, rows := range data{
		response = append(response, ToAdminUserResponse(rows))
	}
	return response, nil
}