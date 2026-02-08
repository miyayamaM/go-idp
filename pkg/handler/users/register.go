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

type UsersRegisterHandlerInterface interface {
	Execute(c *echo.Context) error
}

// UsersRegisterHandler は usecase を注入した Echo の HandlerFunc を返す（Wire 用）
type UsersRegisterHandler struct {
	RegisterUsecase users.UsersRegisterUsecase
}

func NewUsersRegisterHandler(registerUsecase users.UsersRegisterUsecase) UsersRegisterHandlerInterface {
	return &UsersRegisterHandler{RegisterUsecase: registerUsecase}
}

func (h UsersRegisterHandler) Execute(c *echo.Context) error {
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

	_ = h.RegisterUsecase.Execute()

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("register %s", request.Email),
	})
}
