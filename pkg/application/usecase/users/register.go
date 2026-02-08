package users

import users "github.com/miyayamamasaru/go-idp/pkg/domain/interfaces"

type UsersRegisterUsecase interface {
	Execute() error
}

type usersRegisterUsecase struct {
	repo users.UsersRepositoryInterface
}

func (u *usersRegisterUsecase) Execute() error {
	err := u.repo.Save()
	if err != nil {
		return err
	}
	return nil
}

func NewUsersRegisterUsecase(usersRepository users.UsersRepositoryInterface) UsersRegisterUsecase {
	return &usersRegisterUsecase{repo: usersRepository}
}
