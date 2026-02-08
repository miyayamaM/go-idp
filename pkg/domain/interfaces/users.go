package users

type UsersRepositoryInterface interface {
	Save() error
	FindByEmail(email string) error
}
