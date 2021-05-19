package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Any("get", info_get)
	route.Any("list", info_list)
}

func info_get(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	data := MinerModel.Api_find(address)
	if len(data) > 0 {
		data["all_power"] = SystemParamModel.Api_value("RawBytePower")
		data["blocks_num"] = BlocksModel.Api_count_byMiner(address)
		if data["blocks_num"] == nil {
			data["blocks_num"] = 0
		}
		data["reward"] = 0
		data["sector_size"] = "512Mib"
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func info_list(c *gin.Context) {
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit := 20
	power := SystemParamModel.Api_value("RawBytePower")
	datas := MinerModel.Api_select(page, limit)
	for i, data := range datas {
		data["percent"] = Calc.Any2Float64(data["RawBytePower"]) + 0.01/Calc.Any2Float64(power) + 0.1
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}
