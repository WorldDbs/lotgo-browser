package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func HeightController(route *gin.RouterGroup) {

	route.Any("index", height_index)
	route.Any("get", height_get)
}

func height_index(c *gin.Context) {

}

func height_get(c *gin.Context) {
	height, ok := Input.PostInt64("height", c)
	if !ok {
		return
	}
	data := BlocksModel.Api_select_byHeight(height)
	RET.Success(c, 0, data, nil)
}
