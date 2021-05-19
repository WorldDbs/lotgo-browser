package Filecoin

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/config/app_conf"
	"main.go/tuuz/ConfFile"
	"main.go/tuuz/Net"
)

type _MinerGetBaseInfo struct {
	Jsonrpc string                  `json:"jsonrpc"`
	Result  MinerGetBaseInfo_Struct `json:"result"`
	ID      int                     `json:"id"`
}

type MinerGetBaseInfo_Struct struct {
	MinerPower   string `json:"MinerPower"`
	NetworkPower string `json:"NetworkPower"`
	Sectors      []struct {
		SealProof    int `json:"SealProof"`
		SectorNumber int `json:"SectorNumber"`
		SealedCID    struct {
			NAMING_FAILED string `json:"/"`
		} `json:"SealedCID"`
	} `json:"Sectors"`
	WorkerKey       string `json:"WorkerKey"`
	SectorSize      int    `json:"SectorSize"`
	PrevBeaconEntry struct {
		Round int    `json:"Round"`
		Data  string `json:"Data"`
	} `json:"PrevBeaconEntry"`
	BeaconEntries     interface{} `json:"BeaconEntries"`
	EligibleForMining bool        `json:"EligibleForMining"`
}

func MinerGetBaseInfo(address interface{}) (MinerGetBaseInfo_Struct, error) {
	data := BlocksModel.Api_find_new()
	var jr jsonrpc
	jr.Jsonrpc = "2.0"
	jr.Method = "Filecoin.MinerGetBaseInfo"
	jr.ID = 1
	jr.Params = []interface{}{address, data["height"], []interface{}{
		map[string]interface{}{
			"/": data["cid"],
		},
	}}
	get := map[string]interface{}{
		"token": ConfFile.Load("token", "token"),
	}
	ret, err := Net.PostRaw(app_conf.Address, get, jr)
	fmt.Println(ret, err)
	if err != nil {
		fmt.Println(err)
		return MinerGetBaseInfo_Struct{}, err
	} else {
		var data _MinerGetBaseInfo
		jsr := jsoniter.ConfigCompatibleWithStandardLibrary
		err := jsr.UnmarshalFromString(ret, &data)
		if err != nil {
			return MinerGetBaseInfo_Struct{}, err
		}
		return data.Result, nil
	}
}
