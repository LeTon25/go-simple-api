package repo

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct{}

// GetUserByEmail implements IUserRepo.
func (u *userRepo) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepo {
	return &userRepo{}
}
