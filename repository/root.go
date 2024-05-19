package repository

import (
	"sync"
)

// 3티어 아키텍처
// 이 프로젝트에서는 DB를 사용하지 않으므로 서버의 메모리를 통해 관리 되는 값들로 활용할 예정

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	// repository, DB 대한 값 가져야함 (여기서는 일단 제외)
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})
	return &Repository{}
}
