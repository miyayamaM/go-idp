package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/miyayamamasaru/go-idp/pkg/handler"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/health-check", handler.HealthCheck)

	return e
}
