package service

import "GOLANGSERVER/repository"

type User struct {
	UserRepository *repository.UserRepository
}

func netUserService(userRepository *repository.UserRepository) *User {
	return &User{}
}
