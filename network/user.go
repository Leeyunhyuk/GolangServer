package network

import "sync"

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	//service
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {

	})

	return userRouterInstance
}
