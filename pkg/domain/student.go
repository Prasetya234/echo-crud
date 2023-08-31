package domain

import "crud_echo/pkg/dto"

type Student struct {
	Id         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

type StudentRepository interface {
	CreateStudent(req Student) error
	UpdateStudent(id int, req Student) error
	GetStudent() ([]Student, error)
	GetStudentById(id int) (Student, error)
	DeleteStudentById(id int) error
}

type StudentUsecase interface {
	CreateStudent(req dto.StudentDTO) error
	UpdateStudent(id int, req dto.StudentDTO) error
	GetStudent() ([]Student, error)
	GetStudentById(id int) (Student, error)
	DeleteStudentById(id int) error
}
