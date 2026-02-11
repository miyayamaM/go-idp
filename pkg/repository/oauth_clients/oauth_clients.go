package oauth_clients

import (
	"context"
	"database/sql"
	"errors"

	domain "github.com/miyayamamasaru/go-idp/pkg/domain/interfaces"
	models "github.com/miyayamamasaru/go-idp/pkg/repository/models"
	"github.com/uptrace/bun"
)

type OauthClientRepository struct {
	db *bun.DB
}

func NewOauthClientRepository(db *bun.DB) domain.OauthClientRepositoryInterface {
	return &OauthClientRepository{db: db}
}

func (r *OauthClientRepository) FindByClientId(clientID string) (*models.OauthClient, error) {
	var client models.OauthClient
	err := r.db.NewSelect().
		Model(&client).
		Where("client_id = ?::uuid", clientID).
		Scan(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &client, nil
}
