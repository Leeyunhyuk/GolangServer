package service

import (
	"GOLANGSERVER/repository"
	"GOLANGSERVER/types"
)

type User struct {
	UserRepository *repository.UserRepository
}

func netUserService(userRepository *repository.UserRepository) *User {
	return &User{
		UserRepository: userRepository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.UserRepository.Create(newUser)
}

// 실제 업데이트 할 때는 특정 값만 넘겨 줄 것.
func (u *User) Update(beforeUser *types.User, updateUser *types.User) error {
	return u.UserRepository.Update(beforeUser, updateUser)
}
func (u *User) Delete(user *types.User) error {
	return u.UserRepository.Delete(user)
}
func (u *User) Get() []*types.User {
	return u.UserRepository.Get()
}
