package usecase

import (
	"crud_echo/pkg/domain"
	"crud_echo/pkg/dto"
	"crud_echo/shared/util"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		UserRepository: repository,
	}
}

func (u UserUsecase) CreateUser(req dto.UserDTO) error {
	var user domain.User
	mapstructure.Decode(req, &user)
	if _, err := u.UserRepository.FindByEmail(user.Email); err == nil {
		return errors.New("Email already exists")
	}
	user.Password = util.EncryptPassword(user.Password)
	return u.UserRepository.CreateUser(user)
}

func (u UserUsecase) UpdateUser(id int, req dto.UserDTO) error {
	var user domain.User
	mapstructure.Decode(req, &user)
	if _, err := u.UserRepository.FindByEmail(user.Email); err == nil {
		return errors.New("Email already exists")
	}
	return u.UserRepository.UpdateUser(id, user)
}

func (u UserUsecase) GetUser() ([]domain.User, error) {
	return u.UserRepository.GetUser()
}

func (u UserUsecase) GetUserById(id int) (domain.User, error) {
	return u.UserRepository.GetUserById(id)
}

func (u UserUsecase) DeleteUserById(id int) error {
	return u.UserRepository.DeleteUserById(id)
}

func (u UserUsecase) UserLogin(req dto.LoginRequest) (interface{}, error) {
	var loginResponse dto.LoginResponse
	user, err := u.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email not found")
	}
	passwordValid := util.DecryptPassword(user.Password)
	if passwordValid != req.Password {
		return nil, errors.New("bad credential")
	}
	token, err := util.CreateJwtToken(user)
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(user, &loginResponse)
	loginResponse.Token = token
	return loginResponse, nil
}
