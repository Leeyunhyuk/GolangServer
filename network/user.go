package network

import (
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
}

func newUserRouter(router *Network) *userRouter {
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
	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다.")
	/*
		u.router.okResponse(c, &types.UserResponse{
			ApiResponse: &types.ApiResponse{
				Result:      1,
				Description: "성공입니다.",
			},
			User: nil,
		})
	*/
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
		User:        nil,
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다.")

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다.")

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("success", 1),
	})
}
