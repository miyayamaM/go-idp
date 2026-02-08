package users

import users "github.com/miyayamamasaru/go-idp/pkg/domain/interfaces"

type UsersRegisterUsecase interface {
	Execute(email string, password string) error
}

type usersRegisterUsecase struct {
	repo users.UsersRepositoryInterface
}

func (u *usersRegisterUsecase) Execute(email string, password string) error {
	err := u.repo.Save(email, password)
	if err != nil {
		return err
	}
	return nil
}

func NewUsersRegisterUsecase(usersRepository users.UsersRepositoryInterface) UsersRegisterUsecase {
	return &usersRegisterUsecase{repo: usersRepository}
}
