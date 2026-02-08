package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/miyayamamasaru/go-idp/pkg/application/usecase/users"
)

type UsersRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UsersRegisterHandler は usecase を注入した Echo の HandlerFunc を返す（Wire 用）
func UsersRegisterHandler(registerUsecase users.UsersRegisterUsecase) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return UsersRegister(c, registerUsecase)
	}
}

func UsersRegister(c *echo.Context, registerUsecase users.UsersRegisterUsecase) error {
	var request UsersRegisterRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	if request.Email == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Email and password are required",
		})
	}

	registerUsecase.Execute()

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("register %s", request.Email),
	})
}
