package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/controller"
)

func IndexRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.IndexController(route.Group("/index"))
	controller.LoginController(route.Group("/login"))
	controller.BlockController(route.Group("block"))
	controller.GasController(route.Group("gas"))
	controller.MinerController(route.Group("miner"))
	controller.HeightController(route.Group("height"))
	controller.MessageController(route.Group("message"))
}
