package dto

import validation "github.com/go-ozzo/ozzo-validation"

type UserDTO struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Token    string `json:"token"`
}

func (s UserDTO) Validation() error {
	err := validation.ValidateStruct(&s,
		validation.Field(&s.FullName, validation.Required),
		validation.Field(&s.Email, validation.Required),
		validation.Field(&s.Password, validation.Required),
		validation.Field(&s.Address, validation.Required),
	)
	if err != nil {
		return err
	}
	return nil
}
