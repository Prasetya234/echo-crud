package domain

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type UserRepository interface {
}

type UserUsecase interface {
}
