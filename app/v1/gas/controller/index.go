package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/MessageBlsModel"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("list", index_list)
	route.Any("all", index_list_all)
}

func index_list_all(c *gin.Context) {
	datas := MessageBlsModel.Api_select_gas_groupMethod()
	RET.Success(c, 0, datas[0], nil)
}

func index_list(c *gin.Context) {
	datas := MessageBlsModel.Api_select_gas_all()
	RET.Success(c, 0, datas, nil)
}
