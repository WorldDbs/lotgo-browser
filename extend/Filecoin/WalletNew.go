package Filecoin

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/wallet/model/WalletUserModel"
	"main.go/app/v1/wallet/struct/WalletNewStruct"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func WalletNew(password interface{}) (string, string, string, error) {
	token := Calc.GenerateToken()
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.WalletNew"
	jr.ID = 1
	jr.Params = []interface{}{"bls"}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	fmt.Println(ret, err)
	if err != nil {
		fmt.Println(err)
		return "", "", token, err
	} else {
		var data WalletNewStruct.WalletNew
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return "", "", token, err
		}
		bls := data.Result
		pvk, err := WalletExport(bls)
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
			return "", "", token, err
		}
		if WalletUserModel.Api_insert(bls, pvk, password, token) {
			return bls, pvk, token, nil
		} else {
			return "", "", token, errors.New("数据库错误")
		}
	}
}
