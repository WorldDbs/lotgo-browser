package Filecoin

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

type _result struct {
	Jsonrpc string                 `json:"jsonrpc"`
	Result  StateMinerPower_struct `json:"result"`
	ID      int                    `json:"id"`
}
type StateMinerPower_struct struct {
	MinerPower struct {
		RawBytePower    string `json:"RawBytePower"`
		QualityAdjPower string `json:"QualityAdjPower"`
	} `json:"MinerPower"`
	TotalPower struct {
		RawBytePower    string `json:"RawBytePower"`
		QualityAdjPower string `json:"QualityAdjPower"`
	} `json:"TotalPower"`
	HasMinPower bool `json:"HasMinPower"`
}

func StateMinerPower(miner interface{}) (StateMinerPower_struct, error) {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.StateMinerPower"
	jr.ID = 1
	jr.Params = []interface{}{miner, nil}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	//fmt.Println(ret, err)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return StateMinerPower_struct{}, err
	} else {
		var data _result
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return StateMinerPower_struct{}, err
		}
		return data.Result, nil
	}
}
