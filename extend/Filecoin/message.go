package Filecoin

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/index/model/MessageBlsModel"
	"main.go/app/v1/index/model/MessageModel"
	"main.go/app/v1/index/model/MessageSecpkModel"
	"main.go/app/v1/index/struct/MessageStruct"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

func ChainGetBlockMessages(cid interface{}, height int64, miner interface{}) error {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.ChainGetBlockMessages"
	jr.ID = 1
	jr.Params = []interface{}{map[string]interface{}{
		"/": cid,
	}}

	ret, err := Net.PostRaw(app_conf.Address, jr)
	//fmt.Println(ret, err)
	if err != nil {
		fmt.Println(err)
	} else {
		var sta MessageStruct.Message
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &sta)
		if err != nil {
			return err
		} else {
			db := tuuz.Db()
			db.Begin()
			var m MessageModel.Interface
			m.Db = db
			var mbls MessageBlsModel.Interface
			mbls.Db = db
			var mkls MessageSecpkModel.Interface
			mkls.Db = db
			data, err := Jsong.Encode(sta.Result)
			if err != nil {
				return err
			}
			m.Api_insert(cid, height, data, len(sta.Result.BlsMessages), len(sta.Result.SecpkMessages), len(sta.Result.Cids))
			for _, message := range sta.Result.BlsMessages {
				mbls.Api_insert(cid, height, miner, message.CID.NAMING_FAILED, message.Version, message.From, message.To, message.Nonce,
					message.Value, message.GasLimit,
					message.GasFeeCap, message.GasPremium,
					message.Method,
					message.Params)
			}
			for _, secp := range sta.Result.SecpkMessages {
				message := secp.Message
				signature, err := Jsong.Encode(secp.Signature)
				if err != nil {
					Log.Crrs(err, tuuz.FUNCTION_ALL())
					return err
				}
				mkls.Api_insert(cid, height, message.CID.NAMING_FAILED, message.Version, message.From, message.To, message.Nonce, message.Value, message.GasLimit, message.GasFeeCap, message.GasPremium,
					message.Method,
					message.Params, signature, secp.CID.NAMING_FAILED)
			}
			db.Commit()
		}
	}
	return nil
}
