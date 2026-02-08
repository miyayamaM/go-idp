package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/miyayamamasaru/go-idp/pkg/application/usecase/users"
	"github.com/miyayamamasaru/go-idp/pkg/handler"
	handler_users "github.com/miyayamamasaru/go-idp/pkg/handler/users"
)

func NewRouter(registerUsecase users.UsersRegisterUsecase) *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/health-check", handler.HealthCheck)
	e.POST("/users/register", handler_users.UsersRegisterHandler(registerUsecase))

	return e
}
