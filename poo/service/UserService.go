package service

import (
	"foohandler/poo/model"
)

type IUserRepository interface {
	InsertUser(user *model.User)
}

type userService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) InsertUser(user *model.User) {
	u.userRepository.InsertUser(user)
}
