package service

import (
	"errors"
	"net/mail"
	"simple-product-api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{
	Repo models.UserRepository
}

//constructors
func NewUserService(repo models.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func ToAdminUserResponse(user *models.User) (*models.AdminUserResponse){
	return &models.AdminUserResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
	}
}

func ToUserResponse(user *models.User) (*models.UserResponse){
	return &models.UserResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
	}
}

func Validate (req *models.UserRequest) (error){
	//validasi nama
	if len(req.Name) < 3 {
		return errors.New("Name must be at least 3 Characters long!")
	}
	if req.Name == "" {
		return errors.New("Name must not be empty!")
	}

	//validasi email
	err := validateEmail(req.Email)
	if err != nil {
		return errors.New("Invalid Email!")
	}

	//validasi password, nerima 8
	if len(req.Password) <= 7 {
		return errors.New("Password must be at least 8 Characters long!")
	}
	if req.Name == "" {
		return errors.New("Password must not be empty!")
	}
	//berarti aman
	return nil
}

//mungkin perlu implement unique email
func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}


func (us *UserService) GetAllUsers()([]*models.AdminUserResponse, error){
	var response []*models.AdminUserResponse
	
	data, err := us.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, rows := range data{
		response = append(response, ToAdminUserResponse(rows))
	}
	return response, nil
}

func (us *UserService) AdminGetUserbyId(id int) (*models.AdminUserResponse, error) {
	
	data, err := us.Repo.GetUserbyId(id)
	if err != nil {
		return nil, err
	}

	return ToAdminUserResponse(data), nil
}

func (us *UserService) GetUserbyId(id int) (*models.UserResponse, error) {
	
	data, err := us.Repo.GetUserbyId(id)
	if err != nil {
		return nil, err
	}

	return ToUserResponse(data), nil
}

func (us *UserService) Register(req *models.UserRequest) (*models.UserResponse, error) {
	
	//validasi
	err := Validate(req)
	if err != nil {
		return nil, err
	}

	//panggil fungsi hash password, hasilnya diset sebagai password data
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var data = &models.User{
		Name: req.Name,
		Password: string(hashedPassword),
		Email: req.Email,
		Role: models.RoleUser,
	}

	data, err = us.Repo.Register(data)
	if err != nil {
		return nil, err
	}
	
	return ToUserResponse(data), nil
}

func (ur *UserService) Login(req *models.LoginRequest) (*models.User, error){
	//Alur: get user from repo, bandingin password
	data, err := ur.Repo.Login(req.Email)
	if err != nil {
		return nil, err
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	
	//placeholder untuk generate jwt token

	//return data dengan jwt
	return data, nil
}

func (us *UserService) UpdateUser(id int, req *models.UserRequest) (*models.UserResponse, error) {
	
	//validasi
	err := Validate(req)
	if err != nil {
		return nil, err
	}

	//panggil fungsi hash password, hasilnya diset sebagai password data
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var data = &models.User{
		Name: req.Name,
		Password: string(hashedPassword),
		Email: req.Email,
	}

	data, err = us.Repo.UpdateUser(id, data)
	if err != nil {
		return nil, err
	}

	//if userRole = admin, toAdminUserResponse
	return ToUserResponse(data), nil
}

func (us *UserService) DeleteUser(id int) (*models.AdminUserResponse, error) {
	
	data, err := us.Repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return ToAdminUserResponse(data), nil
}
