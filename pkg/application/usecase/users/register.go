package users

import (
	users "github.com/miyayamamasaru/go-idp/pkg/domain/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type UsersRegisterUsecase interface {
	Execute(email string, password string) error
}

type usersRegisterUsecase struct {
	repo users.UsersRepositoryInterface
}

func (u *usersRegisterUsecase) Execute(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return u.repo.Save(email, string(hashedPassword))
}

func NewUsersRegisterUsecase(usersRepository users.UsersRepositoryInterface) UsersRegisterUsecase {
	return &usersRegisterUsecase{repo: usersRepository}
}
