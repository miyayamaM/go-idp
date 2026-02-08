//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v5"
	users_usecase "github.com/miyayamamasaru/go-idp/pkg/application/usecase/users"
	"github.com/miyayamamasaru/go-idp/pkg/repository/users"
)

var SuperSet = wire.NewSet(
	ProvideDB,
	users.NewUsersRepository,
	users_usecase.NewUsersRegisterUsecase,
)

func InitializeRouter() (*echo.Echo, error) {
	wire.Build(SuperSet, NewRouter)
	return nil, nil
}
