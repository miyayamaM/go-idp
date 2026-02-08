package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/miyayamamasaru/go-idp/pkg/handler"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	// InitializeHander は wire_gen.go で生成された実装を使用（通常ビルド時は !wireinject で wire_gen が含まれる）
	handlers, err := InitializeHander()
	if err != nil {
		panic(err)
	}

	e.GET("/health-check", handler.HealthCheck)
	e.POST("/users/register", func(c *echo.Context) error {
		return handlers.UsersRegisterHandler.Execute(c)
	})

	return e
}
