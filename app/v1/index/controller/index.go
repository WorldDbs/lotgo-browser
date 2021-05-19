package controller

import "github.com/gin-gonic/gin"

func IndexController(route *gin.RouterGroup) {

	route.Any("preview", block_preview)
}
