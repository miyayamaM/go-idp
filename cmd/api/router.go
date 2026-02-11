package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/miyayamamasaru/go-idp/pkg/handler"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	handlers, err := InitializeHander()
	if err != nil {
		panic(err)
	}

	e.GET("/health-check", handler.HealthCheck)
	e.POST("/users/register", func(c *echo.Context) error {
		return handlers.UsersRegisterHandler.Execute(c)
	})
	e.GET("/oauth/authorize", func(c *echo.Context) error {
		return handlers.OauthAuthorizeHandler.Execute(c)
	})

	return e
}
