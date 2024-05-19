package network

import (
	"GOLANGSERVER/service"

	"github.com/gin-gonic/gin"
)

//api에 대한 router 설정을 하고 root.go를 통해 api 관리를 root만 cmd에 넣어줘서 관리

type Network struct {
	engin *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r, service.User)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}
