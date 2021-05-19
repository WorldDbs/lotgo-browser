package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/app/v1/index/model/MessageBlsModel"
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
	timestamp := int64(0)
	datas := BlocksModel.Api_select_byHeight(height)
	for i, data := range datas {
		timestamp = data["timestamp"].(int64)
		data["message_count"] = MessageBlsModel.Api_count_byBelongCid(data["cid"])
		data["reward"] = MessageBlsModel.Api_sum_value(data["cid"])
		if data["reward"] == nil {
			data["reward"] = 0
		}
		datas[i] = data
	}
	RET.Success(c, 0, datas, timestamp)
}
