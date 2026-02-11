//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	users_usecase "github.com/miyayamamasaru/go-idp/pkg/application/usecase/users"
	handler_oauth "github.com/miyayamamasaru/go-idp/pkg/handler/oauth"
	handler_users "github.com/miyayamamasaru/go-idp/pkg/handler/users"
	oauth_clients_repository "github.com/miyayamamasaru/go-idp/pkg/repository/oauth_clients"
	users_repository "github.com/miyayamamasaru/go-idp/pkg/repository/users"
)

var SuperSet = wire.NewSet(
	ProvideDB,
	ProvideBunDB,
)

// repository
var repositorySet = wire.NewSet(
	users_repository.NewUsersRepository,
	oauth_clients_repository.NewOauthClientRepository,
)

// usecase
var usecaseSet = wire.NewSet(
	users_usecase.NewUsersRegisterUsecase,
)

// handler
var handlerSet = wire.NewSet(
	handler_users.NewUsersRegisterHandler,
	handler_oauth.NewAuthorizeHandler,
)

type HandlerSet struct {
	UsersRegisterHandler  handler_users.UsersRegisterHandlerInterface
	OauthAuthorizeHandler handler_oauth.AuthorizeHandlerInterface
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
