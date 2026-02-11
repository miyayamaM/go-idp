package interfaces

type UsersRepositoryInterface interface {
	Save(email string, password string) error
	FindByEmail(email string) error
}
