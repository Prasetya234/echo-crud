package router

import (
	"crud_echo/pkg/controller"
	"crud_echo/pkg/repository"
	"crud_echo/pkg/usecase"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := &controller.UserController{
		UserUsecase: uu,
	}
	e.POST("/register", uc.CreateUser)
	e.POST("/login", uc.Login)
	g.GET("/user", uc.GetUser)
	g.GET("/user/:id", uc.GetUserById)
	g.PUT("/user/:id", uc.UpdateUser)
	g.DELETE("/user/:id", uc.DeleteUser)
}
