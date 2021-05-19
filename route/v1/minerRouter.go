package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/miner/controller"
)

func MinerRouter(route *gin.RouterGroup) {

	controller.IndexController(route.Group("index"))
	controller.InfoController(route.Group("info"))
	controller.PlegController(route.Group("pleg"))
	controller.StaticController(route.Group("static"))

}
