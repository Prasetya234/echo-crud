package controller

import (
	"crud_echo/pkg/domain"
	"crud_echo/pkg/dto"
	"crud_echo/shared/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) GetUser(c echo.Context) error {
	resp, err := uc.UserUsecase.GetUser()
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

func (uc *UserController) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := uc.UserUsecase.GetUserById(id)

	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "id not found", nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var request dto.UserDTO
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := uc.UserUsecase.CreateUser(request); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	var request dto.UserDTO

	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err2 := uc.UserUsecase.UpdateUser(id, request); err2 != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err2.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := uc.UserUsecase.GetUserById(id)
	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "id not found", nil)
	}
	if err := uc.UserUsecase.DeleteUserById(id); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}
func (uc *UserController) Login(c echo.Context) error {
	var request dto.LoginRequest
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}
	resp, err := uc.UserUsecase.UserLogin(request)
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}
