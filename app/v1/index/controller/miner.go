package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func MinerController(route *gin.RouterGroup) {
	route.Any("list", miner_list)
}

func miner_list(c *gin.Context) {
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit := 20
	datas := MinerModel.Api_select(page, limit)
	for i, data := range datas {
		data["tag"] = data["miner"]
		data["amount"] = 0
		data["percent"] = 0.1
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}
