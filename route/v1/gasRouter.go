package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/gas/controller"
)

func GasRouter(route *gin.RouterGroup) {
	controller.IndexController(route.Group("index"))
	controller.ChartController(route.Group("chart"))
}
