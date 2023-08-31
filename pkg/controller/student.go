package controller

import (
	"crud_echo/pkg/domain"
	"crud_echo/pkg/dto"
	"crud_echo/shared/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type StudentController struct {
	StudentUsecase domain.StudentUsecase
}

func (sc *StudentController) GetStudent(c echo.Context) error {
	resp, err := sc.StudentUsecase.GetStudent()
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

func (sc *StudentController) GetStudentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := sc.StudentUsecase.GetStudentById(id)

	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "student id not found", nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

func (sc *StudentController) CreateStudent(c echo.Context) error {
	var request dto.StudentDTO
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := sc.StudentUsecase.CreateStudent(request); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}

func (sc *StudentController) UpdateStudent(c echo.Context) error {
	var request dto.StudentDTO

	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err2 := sc.StudentUsecase.UpdateStudent(id, request); err2 != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err2.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}

func (sc *StudentController) DeleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := sc.StudentUsecase.GetStudentById(id)
	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "id not found", nil)
	}
	if err := sc.StudentUsecase.DeleteStudentById(id); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}
