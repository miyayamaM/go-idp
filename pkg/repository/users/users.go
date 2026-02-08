package users

import (
	"database/sql"

	domain "github.com/miyayamamasaru/go-idp/pkg/domain/interfaces"
)

type UsersRepository struct {
	db *sql.DB
}

// NewUsersRepository は domain の UsersRepositoryInterface を実装したリポジトリを返す
func NewUsersRepository(db *sql.DB) domain.UsersRepositoryInterface {
	return &UsersRepository{db: db}
}

// Save は domain/interfaces の UsersRepositoryInterface の実装
func (r *UsersRepository) Save() error {
	// TODO: 実装
	return nil
}

// FindByEmail は domain/interfaces の UsersRepositoryInterface の実装
func (r *UsersRepository) FindByEmail(email string) error {
	// TODO: 実装
	return nil
}
