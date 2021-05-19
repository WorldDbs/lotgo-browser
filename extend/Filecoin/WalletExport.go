package Filecoin

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/wallet/struct/WalletExportSturct"
	"main.go/config/app_conf"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Net"
)

func WalletExport(bls_address interface{}) (string, error) {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.WalletExport"
	jr.ID = 1
	jr.Params = []interface{}{bls_address}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	if err != nil {
		return "", err
	} else {
		var data WalletExportSturct.WalletExport
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return "", err
		}
		return data.Result.PrivateKey, nil
	}
}
