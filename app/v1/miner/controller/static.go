package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/app/v1/miner/model/MinerPowerLogModel"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"time"
)

func StaticController(route *gin.RouterGroup) {
	route.Any("24", static_24)
}

func static_24(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	timestamp := time.Now().Unix()
	start_data := MinerPowerLogModel.Api_find_one(address, ">", "asc", timestamp-86400)
	if len(start_data) < 1 {
		start_data = MinerPowerLogModel.Api_find_one(address, "<", "asc", timestamp)
	}
	end_data := MinerPowerLogModel.Api_find_one(address, "<", "desc", timestamp)
	if end_data["amount"] == "" {
		end_data["amount"] = 0
	}
	if end_data["MinerPower"] == "" {
		end_data["MinerPower"] = 0
	}
	if end_data["NetworkPower"] == "" {
		end_data["NetworkPower"] = 0
	}
	if end_data["MinerPower"] == "" {
		end_data["MinerPower"] = 0
	}
	amount := end_data["amount"]
	rewardPerByte := Calc.Any2Float64(amount) / (Calc.Any2Float64(end_data["MinerPower"]) + 1)
	MinerPowerLogModel.Api_select(address, timestamp-86400, timestamp)
	blocks_num := BlocksModel.Api_count_byMiner(address)
	miner := MinerModel.Api_find(address)
	data := map[string]interface{}{
		"blocksMined":            blocks_num,
		"durationPercentage":     0.1,
		"equivalentMiners":       miner["WorkerKey"],
		"luckyValue":             "N/A",
		"networkTotalRewards":    amount,
		"qualityAdjPowerDelta":   end_data["MinerPower"],
		"qualityAdjPowerGrowth":  Calc.Any2String(Calc.Any2Float64(end_data["NetworkPower"]) - Calc.Any2Float64(start_data["MinerPower"])),
		"rawBytePowerDelta":      end_data["NetworkPower"],
		"rawBytePowerGrowth":     Calc.Any2String(Calc.Any2Float64(end_data["NetworkPower"]) - Calc.Any2Float64(start_data["MinerPower"])),
		"rewardPerByte":          rewardPerByte,
		"totalRewards":           Calc.Any2String(Calc.Any2Float64(end_data["amount"]) - Calc.Any2Float64(start_data["amount"])),
		"weightedBlocksMined":    "0",
		"windowedPoStFeePerByte": "0",
	}
	RET.Success(c, 0, data, nil)
}
