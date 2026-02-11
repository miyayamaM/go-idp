//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	users_usecase "github.com/miyayamamasaru/go-idp/pkg/application/usecase/users"
	handler_users "github.com/miyayamamasaru/go-idp/pkg/handler/users"
	users_repository "github.com/miyayamamasaru/go-idp/pkg/repository/users"
)

var SuperSet = wire.NewSet(
	ProvideDB,
	ProvideBunDB,
)

// repository
var repositorySet = wire.NewSet(
	users_repository.NewUsersRepository,
)

// usecase
var usecaseSet = wire.NewSet(
	users_usecase.NewUsersRegisterUsecase,
)

// handler
var handlerSet = wire.NewSet(
	handler_users.NewUsersRegisterHandler,
)

type HandlerSet struct {
	UsersRegisterHandler handler_users.UsersRegisterHandlerInterface
}

func InitializeHander() (*HandlerSet, error) {
	wire.Build(
		SuperSet,
		repositorySet,
		usecaseSet,
		handlerSet,
		wire.Struct(new(HandlerSet), "*"),
	)
	return nil, nil
}
