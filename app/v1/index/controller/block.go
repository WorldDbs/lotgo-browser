package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/app/v1/index/model/HeightModel"
	"main.go/app/v1/index/model/MessageBlsModel"
	"main.go/app/v1/index/model/MessageModel"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/app/v1/miner/model/MinerPowerLogModel"
	"main.go/app/v1/wallet/model/WalletUserModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"math"
	"time"
)

func BlockController(route *gin.RouterGroup) {

	route.Use(BaseController.CommonController(), gin.Recovery())

	route.Any("preview", block_preview)
	route.Any("list", block_index)
	route.Any("new", block_new)
	route.Any("message", block_message)
	route.Any("get", block_get)

}

func block_preview(c *gin.Context) {
	miner_power := Calc.Any2Float64(MinerModel.Api_sum_QualityAdjPower()) / math.Pow(1024, 3)
	block_height := BlocksModel.Api_find_new()
	gas_price := SystemParamModel.Api_value("gas")
	miners := MinerModel.Api_count()
	block_last2 := BlocksModel.Api_find(block_height["height"].(int64) - 1)
	message := MessageModel.Api_find_height(block_height["height"])
	num := MessageModel.Api_count_24()
	//num_24 := BlocksModel.Api_count_24()
	//blsmessage := MessageBlsModel.Api_select_byMethod([]interface{}{4})
	bls_message_all := MessageBlsModel.Api_select_byMethod([]interface{}{0, 1, 2, 3, 4, 5, 6, 7})
	zhiya := MessageBlsModel.Api_select_byMethod([]interface{}{2})
	if zhiya == nil {
		zhiya = 0
	}
	wallet_num := WalletUserModel.Api_count()
	rand := Calc.Rand(2000000, 3000000)
	rand2 := Calc.Rand(5000, 6000)
	//pleg := SystemParamModel.Api_value("pleg")
	timestamp := time.Now().Unix()

	destory := Calc.Any2Float64(SystemParamModel.Api_value("FilBurnt")) / 1000000000000000000
	lock_amount := Calc.Any2Float64(SystemParamModel.Api_value("FilLocked")) / 1000000000000000000
	mined_amount := Calc.Any2Float64(SystemParamModel.Api_value("FilMined"))
	start_amount := MinerPowerLogModel.Api_sum_amount_all(timestamp - 86400 - 600)
	if start_amount == nil {
		start_amount = 0
	}
	end_amount := MinerPowerLogModel.Api_sum_amount_all(timestamp - 600)

	h_reward := Calc.Round((Calc.Any2Float64(end_amount)-Calc.Any2Float64(start_amount))/1000000000000000000, 4)
	mined_amount = mined_amount / 1000000000000000000
	circulating := SystemParamModel.Api_value("FilCirculating")
	data := map[string]interface{}{
		"block_height":         block_height["height"].(int64),
		"last_block_time":      block_height["timestamp"].(int64) - time.Now().Unix(),
		"miner_power":          Calc.Any2String(miner_power),
		"active_miner":         miners,
		"block_reward":         Calc.Round(float64(rand)/1000000.0, 2),
		"24h_reward_per_tib":   0.006,
		"24h_reward":           Calc.Any2String(h_reward),
		"current_pleg":         Calc.Round(Calc.Any2Float64(zhiya)/1000000000000000000, 2),
		"pleg_number":          Calc.Round(Calc.Any2Float64(zhiya)/1000000000000000000, 2),
		"pleg_number_per_tib":  float64(rand2) / 10000,
		"24h_message_number":   num,
		"block_time":           Calc.Any2Int64(block_height["timestamp"]) - Calc.Any2Int64(block_last2["timestamp"]),
		"block_num_per_height": message["num_cid"],
		"miner_power_per_tib":  8.9,
		"price":                "N/A",
		"gas_price":            Calc.Round(Calc.Any2Float64(gas_price)/100000, 4),
		"wallet_num":           wallet_num,
		"bls_amount_all":       Calc.Round(Calc.Any2Float64(bls_message_all)/1000000000000000000, 2),
		"all_num":              "200 Million ",
		"destory":              Calc.Round(destory, 5),
		"lock_amount":          Calc.Round(lock_amount, 4),
		"mined_amount":         Calc.Round(mined_amount, 4),
		"circulating_real":     Calc.Round(Calc.Any2Float64(circulating)/1000000000000000000, 4),
		"circulating_calc":     Calc.Round(mined_amount-lock_amount-destory, 4),
	}
	RET.Success(c, 0, data, nil)
}

func block_new(c *gin.Context) {
	hs := HeightModel.Api_select(1, 10)
	data := map[int64]interface{}{}
	for _, h := range hs {
		data[h["height"].(int64)] = BlocksModel.Api_select_byHeight(h["height"])
	}
	RET.Success(c, 0, data, nil)
}

func block_index(c *gin.Context) {
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit := 20
	count := HeightModel.Api_count()
	hs := HeightModel.Api_select(page, limit)
	data := map[int64]interface{}{}
	for _, h := range hs {
		blocks := BlocksModel.Api_select_byHeight(h["height"])
		for i, block := range blocks {
			block["reward"] = 0
			block["text"] = "-"
			blocks[i] = block
		}
		data[h["height"].(int64)] = blocks
	}
	RET.Success(c, 0, map[string]interface{}{
		"count": count,
		"page":  count / int64(limit),
		"limit": limit,
		"data":  data,
	}, nil)
}

func block_message(c *gin.Context) {
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit := 20
	count := BlocksModel.Api_count()
	hs := BlocksModel.Api_select(page, limit)
	RET.Success(c, 0, map[string]interface{}{
		"count": count,
		"page":  count / int64(limit),
		"limit": limit,
		"data":  hs,
	}, nil)
}

func block_get(c *gin.Context) {
	cid, ok := Input.Post("cid", c, true)
	if !ok {
		return
	}
	data := BlocksModel.Api_find_byCid(cid)
	if len(data) > 0 {
		data["message_count"] = MessageBlsModel.Api_count_byBelongCid(data["cid"])
		data["reward"] = MessageBlsModel.Api_sum_value(data["cid"])
		if data["reward"] == nil {
			data["reward"] = 0
		}
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
