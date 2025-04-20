package user

import "github.com/hanzohasashi17/farm/internal/user/model"

type UserService struct {
	*UserRepo
}

func NewUserService(userRepo *UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}