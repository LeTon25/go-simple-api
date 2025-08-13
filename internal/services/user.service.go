package services

import (
	"go-simple-api/internal/repo"
	"go-simple-api/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepo
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	//
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorUserAlreadyExists
	}
	return response.ErrorCodeSuccess
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
