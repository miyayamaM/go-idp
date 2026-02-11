package interfaces

import "github.com/miyayamamasaru/go-idp/pkg/repository/models"

type OauthClientRepositoryInterface interface {
	FindByClientId(client_id string) (*models.OauthClient, error)
}
