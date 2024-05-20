package repository

import "GOLANGSERVER/types"

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

// 실제 업데이트 할 때는 특정 값만 넘겨 줄 것.
func (u *UserRepository) Update(beforeUser *types.User, updateUser *types.User) error {
	return nil
}
func (u *UserRepository) Delete(user *types.User) error {
	return nil
}
func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
