package cmd

import (
	"crud_echo/config"
	"crud_echo/pkg/router"
	"crud_echo/shared/db"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {
	e := echo.New()
	g := e.Group("")
	conf := config.GetConfig()
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS512" {
					return nil, fmt.Errorf("token jwt salah format")
				}
				return []byte(conf.SignKey), nil
			}
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, err
			}
			return token, err
		},
	}))
	Apply(e, g, conf)
	e.Logger.Error(e.Start(":5000"))
}

func Apply(e *echo.Echo, g *echo.Group, conf config.Configuration) {
	db := db.NewInstanceDb(conf)
	router.NewStudentRouter(e, g, db)
	router.NewUserRouter(e, g, db)
}
