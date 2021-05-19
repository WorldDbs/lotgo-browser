package Filecoin

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/index/model/PlegModel"
	"main.go/app/v1/wallet/struct/WalletNewStruct"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Net"
	"math"
	"time"
)

func StateMinerInitialPledgeCollateral() {

	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.StateMinerInitialPledgeCollateral"
	jr.ID = 1
	jr.Params = []interface{}{nil, nil, nil}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	//fmt.Println(ret, err)
	if err != nil {
		//fmt.Println(err)

	} else {
		var data WalletNewStruct.WalletNew
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return
		}
		num := SystemParamModel.Api_value("pleg")
		if math.Abs(Calc.Any2Float64(num)-Calc.Any2Float64(data.Result)) > 10000000 {
			return
		}
		SystemParamModel.Api_update("pleg", data.Result)
		PlegModel.Api_insert(data.Result, time.Now().Unix())
	}
}
