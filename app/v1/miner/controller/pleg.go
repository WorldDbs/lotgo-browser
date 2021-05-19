package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/MessageBlsModel"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func PlegController(route *gin.RouterGroup) {
	route.Any("get", pleg_get)
}

func pleg_get(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	miner_info := MinerModel.Api_find(address)
	if len(miner_info) > 0 {
		pleg_amount := MessageBlsModel.Api_sum_To_byMethod(6, address)
		if pleg_amount == nil {
			pleg_amount = 0
		}
		lock_amount := MessageBlsModel.Api_sum_From_byMethod(2, address)
		if lock_amount == nil {
			lock_amount = 0
		}
		after_amount := Calc.Any2Float64(miner_info["amount"]) - Calc.Any2Float64(pleg_amount) - Calc.Any2Float64(lock_amount)

		RET.Success(c, 0, map[string]interface{}{
			"pleg_amount":  Calc.Any2String(pleg_amount),
			"lock_amount":  Calc.Any2String(lock_amount),
			"after_amount": Calc.Any2String(after_amount),
		}, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
