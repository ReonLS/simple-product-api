package models

//Repository Pattern
type UserRepository interface{
	GetAllUsers()([]*User, error)
	GetUserbyId(id int)(*User, error)
	Register(model *User)(*User, error)
	UpdateUser(id int, model *User)(*User, error)
	DeleteUser(id int)(*User, error)
	Login(email string) (*User, error)
}

type User struct{
	Id int
	Name string
	Password string
	Email string
	Role Role
}

type Role string

const (
	RoleUser Role = "User"
	RoleAdmin Role = "Admin"
)

//Update & Register
type UserRequest struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type AdminUserResponse struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role Role `json:"role"`
}