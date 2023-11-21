package main

import (
	"foohandler/poo/model"
	"foohandler/poo/repository"
	"foohandler/poo/service"
)

func main() {
	user := &model.User{
		Name:     "Oséias",
		Password: "1234",
	}

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userService.InsertUser(user)
}
