package models

import "github.com/uptrace/bun"

type OauthClient struct {
	bun.BaseModel `bun:"table:oauth_clients,alias:oc"`

	ID       int64 `bun:",pk,autoincrement"`
	ClientID string
}
