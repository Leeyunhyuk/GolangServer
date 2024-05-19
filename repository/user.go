package repository

import "GOLANGSERVER/types"

type UserRepository struct {
	UserMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{}
}
