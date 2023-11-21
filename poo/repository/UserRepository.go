package repository

import (
	"fmt"
	"foohandler/poo/model"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (u *userRepository) InserUser(user *model.User) {
	//insert on database
	fmt.Println("Insert user on database")
}
