package Filecoin

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/app/v1/index/model/HeightModel"
	"main.go/app/v1/index/struct/HeightSturct"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
	"time"
)

type jsonrpc struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

func Get_block_by_height(height_att int64) {
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.ChainGetTipSetByHeight"
	jr.ID = 1
	jr.Params = []interface{}{height_att, nil}

	ret, err := Net.PostRaw(app_conf.Address, jr)
	//fmt.Println(ret, err)
	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
	} else {
		var height HeightSturct.Height
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &height)
		if err != nil {
			time.Sleep(5 * time.Second)
			return
		}
		if height.Error.Code != 0 {
			return
		}
		if len(HeightModel.Api_find(height.Result.Height)) > 0 {
			Get_block_by_height(height_att + 1)
		} else {
			db := tuuz.Db()
			db.Begin()
			var h HeightModel.Interface
			h.Db = db
			if !h.Api_insert(height.Result.Height, ret) {
				return
			}
			for i, block := range height.Result.Blocks {
				datas, err := Jsong.Encode(block)
				if err != nil {

				} else {
					cid := height.Result.Cids[i]
					if cid.NAMING_FAILED == "" {
						Log.Crrs(errors.New(ret), tuuz.FUNCTION_ALL())
					}
					parents, err := Jsong.Encode(block.Parents)
					if err != nil {
						return
					}
					Messages, err := Jsong.Encode(block.Messages)
					if err != nil {
						return
					}
					var b BlocksModel.Interface
					b.Db = db
					if !b.Api_insert(cid.NAMING_FAILED, height.Result.Height, block.Timestamp, block.Miner, parents, block.ParentWeight, block.Messages.NAMING_FAILED, Messages, datas) {
						fmt.Println(datas)
					}
					err = ChainGetBlockMessages(cid.NAMING_FAILED, height_att)
					if err != nil {
						Log.Crrs(err, tuuz.FUNCTION_ALL())
						return
					}
				}
			}
			db.Commit()
		}
	}
}
