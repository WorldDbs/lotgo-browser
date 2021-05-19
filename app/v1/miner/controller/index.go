package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {

	route.Any("list", info_list)
	route.Any("power", index_minerpower)
}

func index_minerpower(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	data := MinerModel.Api_find(address)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
