package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/extend/Filecoin"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func MinerController(route *gin.RouterGroup) {
	route.Any("list", miner_list)
}

//deceprated in next version
func miner_list(c *gin.Context) {
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit := 20
	datas := MinerModel.Api_select(page, limit)
	for i, data := range datas {
		aaa, _ := Filecoin.WalletBalance(data["miner"])
		data["amount"] = aaa
		data["percent"] = 0.1
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}
