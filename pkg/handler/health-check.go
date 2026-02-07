package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// handler godoc
// @Summary users
// @Description get the users list
// @Accept  json
// @Produce  json
// @Param  service query string true "Service you want to get the users list"
// @Success 200 {object} map[string]string
// @Router /health-check [get]
func HealthCheck(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
