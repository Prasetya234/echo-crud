package usecase

import (
	"crud_echo/pkg/domain"
	"crud_echo/pkg/dto"
	"github.com/mitchellh/mapstructure"
)

type StudentUsecase struct {
	studentRepository domain.StudentRepository
}

func NewStudentUsecase(studentRepository domain.StudentRepository) domain.StudentUsecase {
	return &StudentUsecase{
		studentRepository: studentRepository,
	}
}

func (s StudentUsecase) CreateStudent(req dto.StudentDTO) error {
	var student domain.Student
	if err := mapstructure.Decode(req, &student); err != nil {
		return err
	}
	return s.studentRepository.CreateStudent(student)
}

func (s StudentUsecase) UpdateStudent(id int, req dto.StudentDTO) error {
	var student domain.Student
	if err := mapstructure.Decode(req, &student); err != nil {
		return err
	}
	return s.studentRepository.UpdateStudent(id, student)
}

func (s StudentUsecase) GetStudent() ([]domain.Student, error) {
	return s.studentRepository.GetStudent()
}

func (s StudentUsecase) GetStudentById(id int) (domain.Student, error) {
	return s.studentRepository.GetStudentById(id)
}

func (s StudentUsecase) DeleteStudentById(id int) error {
	return s.studentRepository.DeleteStudentById(id)
}
