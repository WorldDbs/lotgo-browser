package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/wallet/controller"
)

func WalletRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.IndexController(route.Group("index"))
	controller.TransferController(route.Group("transfer"))
	controller.UserController(route.Group("user"))
}
