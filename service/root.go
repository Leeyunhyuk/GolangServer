package service

import (
	"GOLANGSERVER/repository"
	"sync"
)

// Network, repository의 다리 역할,  api를 받아왔을 때 서비스단으로 넘겨 줌
var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	// repository
	repository *repository.Repository

	User *User
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}

		serviceInstance.User = netUserService(rep.User)
	})
	return serviceInstance
}
