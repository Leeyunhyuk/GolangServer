package network

import (
	"GOLANGSERVER/service"
	"GOLANGSERVER/types"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	//service

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)

		// userRouterInstance.router.engin.POST("/", userRouterInstance.create)
		// userRouterInstance.router.engin.PUT("/", userRouterInstance.update)
		// userRouterInstance.router.engin.DELETE("/", userRouterInstance.delete)
		// userRouterInstance.router.engin.GET("/", userRouterInstance.get)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create 입니다.")

	//network단에서는 service단에서만 소통을 하기 때문에 함수만 호출해주면 된다.
	u.userService.Create(nil)

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다.")

	u.userService.Get()
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다.")
	u.userService.Update(nil, nil)
	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다.")

	u.userService.Delete(nil)
	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
	})
}
