package users

import (
	"context"

	domain "github.com/miyayamamasaru/go-idp/pkg/domain/interfaces"
	models "github.com/miyayamamasaru/go-idp/pkg/repository/models"
	"github.com/uptrace/bun"
)

type UsersRepository struct {
	db *bun.DB
}

// NewUsersRepository は domain の UsersRepositoryInterface を実装したリポジトリを返す
func NewUsersRepository(db *bun.DB) domain.UsersRepositoryInterface {
	return &UsersRepository{db: db}
}

// Save は domain/interfaces の UsersRepositoryInterface の実装（Bun で INSERT）
func (r *UsersRepository) Save(email string, password string) error {
	user := &models.User{Email: email, PasswordHash: password}
	_, err := r.db.NewInsert().Model(user).Exec(context.Background())
	return err
}

// FindByEmail は domain/interfaces の UsersRepositoryInterface の実装
func (r *UsersRepository) FindByEmail(email string) error {
	// TODO: 実装
	return nil
}
