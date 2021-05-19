package Filecoin

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/wallet/struct/WalletNewStruct"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func WalletBalance(address interface{}) (string, error) {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.WalletBalance"
	jr.ID = 1
	jr.Params = []interface{}{address}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	//fmt.Println(ret, err)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return "", err
	} else {
		var data WalletNewStruct.WalletNew
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return "", err
		}
		if data.Result == "" {
			return "0", nil
		}
		return data.Result, nil
	}
}
