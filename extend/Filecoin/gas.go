package Filecoin

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/index/model/GasModel"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
	"time"
)

type gas struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}

func GasEstimateGasPremium() {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.GasEstimateGasPremium"
	jr.ID = 1
	jr.Params = []interface{}{42, nil, 9, nil}

	ret, err := Net.PostRaw(app_conf.Address, nil, jr)
	//fmt.Println(ret, err)
	if err != nil {
		fmt.Println(err)
	} else {
		var sta gas
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &sta)
		if err != nil {
			return
		}
		SystemParamModel.Api_update("gas", sta.Result)
		GasModel.Api_insert(sta.Result, time.Now().Unix())
	}
}
