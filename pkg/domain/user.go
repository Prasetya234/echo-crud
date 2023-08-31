package domain

import "crud_echo/pkg/dto"

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type UserRepository interface {
	CreateUser(req User) error
	UpdateUser(id int, req User) error
	GetUser() ([]User, error)
	GetUserById(id int) (User, error)
	DeleteUserById(id int) error
	FindByEmail(email string) (User, error)
}

type UserUsecase interface {
	CreateUser(req dto.UserDTO) error
	UpdateUser(id int, req dto.UserDTO) error
	GetUser() ([]User, error)
	GetUserById(id int) (User, error)
	DeleteUserById(id int) error
	UserLogin(req dto.LoginRequest) (interface{}, error)
}
