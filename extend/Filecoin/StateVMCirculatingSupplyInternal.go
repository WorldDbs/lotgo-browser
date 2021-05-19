package Filecoin

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

type _StateVMCirculatingSupplyInternal struct {
	Jsonrpc string                                  `json:"jsonrpc"`
	Result  StateVMCirculatingSupplyInternal_Struct `json:"result"`
	ID      int                                     `json:"id"`
}

type StateVMCirculatingSupplyInternal_Struct struct {
	FilVested           string `json:"FilVested"`
	FilMined            string `json:"FilMined"`
	FilBurnt            string `json:"FilBurnt"`
	FilLocked           string `json:"FilLocked"`
	FilCirculating      string `json:"FilCirculating"`
	FilReserveDisbursed string `json:"FilReserveDisbursed"`
	Timestamp           int64  `gorose:"timestamp"`
}

func StateVMCirculatingSupplyInternal() (StateVMCirculatingSupplyInternal_Struct, error) {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.StateVMCirculatingSupplyInternal"
	jr.ID = 1
	jr.Params = []interface{}{nil}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	//fmt.Println(ret, err)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return StateVMCirculatingSupplyInternal_Struct{}, err
	} else {
		var data _StateVMCirculatingSupplyInternal
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return StateVMCirculatingSupplyInternal_Struct{}, err
		}
		return data.Result, nil
	}
}
