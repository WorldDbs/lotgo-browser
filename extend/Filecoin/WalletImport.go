package Filecoin

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/wallet/struct/WalletNewStruct"
	"main.go/config/app_conf"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Net"
)

func WalletImport(private_key interface{}) (string, error) {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.WalletImport"
	jr.ID = 1
	jr.Params = []interface{}{map[string]interface{}{
		"Type":       "bls",
		"PrivateKey": private_key,
	}}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	if err != nil {
		return "", err
	} else {
		var data WalletNewStruct.WalletNew
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return "", err
		}
		return data.Result, nil
	}
}
