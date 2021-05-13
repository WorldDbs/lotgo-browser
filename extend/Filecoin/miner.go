package Filecoin

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type state struct {
	Jsonrpc string   `json:"jsonrpc"`
	Result  []string `json:"result"`
	ID      int      `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func StateListMiners() {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.StateListMiners"
	jr.ID = 1
	jr.Params = []interface{}{nil}

	ret, err := Net.PostRaw(app_conf.Address, jr)
	//fmt.Println(ret, err)
	if err != nil {
		fmt.Println(err)
	} else {
		var sta state
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &sta)
		if err != nil {
			return
		}
		if sta.Error.Code != 0 {
			return
		}
		for _, miner := range sta.Result {
			if len(MinerModel.Api_find(miner)) > 0 {

			} else {
				MinerModel.Api_insert(miner)
			}
		}
	}
}
