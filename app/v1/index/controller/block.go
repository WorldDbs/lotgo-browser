package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/app/v1/index/model/HeightModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
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
	block_height := 0
	last_block_time := 0
	gas_price := SystemParamModel.Api_value("gas")
	data := map[string]interface{}{
		"block_height":         block_height,
		"last_block_time":      last_block_time,
		"miner_power":          10000000000,
		"active_miner":         2190,
		"block_reward":         25.89147,
		"24h_reward_per_tib":   0.006,
		"24h_reward":           367565,
		"current_pleg":         0.2560,
		"pleg_number":          70313270,
		"24h_message_number":   100000,
		"block_time":           30.18,
		"block_num_per_height": 4,
		"miner_power_per_tib":  8.9,
		"price":                150,
		"gas_price":            gas_price,
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
	RET.Success(c, 0, data, nil)
}
