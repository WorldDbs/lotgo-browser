package Filecoin

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/wallet/struct/MpoolPushMessageStruct"
	"main.go/config/app_conf"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Net"
)

type Message struct {
	Version    int    `json:"Version"`
	To         string `json:"To"`
	From       string `json:"From"`
	Nonce      int    `json:"Nonce"`
	Value      string `json:"Value"`
	GasLimit   int    `json:"GasLimit"`
	GasFeeCap  string `json:"GasFeeCap"`
	GasPremium string `json:"GasPremium"`
	Method     int    `json:"Method"`
	Params     string `json:"Params"`
	CID        struct {
		NAMING_FAILED string `json:"/"`
	} `json:"CID"`
}

func MpoolPushMessage(from, to, value, GasLimit interface{}) (Message, error) {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.MpoolPushMessage"
	jr.ID = 1
	jr.Params = []interface{}{map[string]interface{}{
		"To":   to,
		"From": from,
		//"Nonce":      0,
		"Value":      value,
		"GasLimit":   GasLimit,
		"GasFeeCap":  "0",
		"GasPremium": "0",
		"Method":     0,
		"Params":     "Ynl0ZSBhcnJheQ==",
	}, map[string]interface{}{
		"MaxFee": "0",
	}}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}

	ret, err := Net.PostRaw(app_conf.Address, get, jr)

	fmt.Println(jr)
	fmt.Println(ret)
	if err != nil {
		return Message{}, err
	} else {
		var data MpoolPushMessageStruct.MpoolPushMessage
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return Message{}, err
		}
		if data.Error.Code != 0 {
			return Message{}, errors.New(data.Error.Message)
		}
		return data.Result.Message, nil
	}
}
