package service

import "github.com/tranvuongduy2003/go-backend/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserInfo() string {
	return us.userRepo.GetUserInfo()
}