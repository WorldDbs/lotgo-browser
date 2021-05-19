package cron

import (
	"main.go/app/v1/index/model/CirculatingModel"
	"main.go/app/v1/index/model/MinerModel"
	"main.go/app/v1/miner/model/MinerPowerLogModel"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/extend/Filecoin"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

func Refresh_60s() {
	refresh_miner_power()
	refresh_miner_amount()
	refresh_state()
	tick := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-tick.C:
			refresh_miner_power()
			refresh_miner_amount()
			refresh_state()
		}
	}
}

func refresh_miner_power() {
	datas := MinerModel.Api_select_all()
	RawBytePower := ""
	QualityAdjPower := ""
	for _, data := range datas {
		miner_power, err := Filecoin.StateMinerPower(data["miner"])

		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			RawBytePower = miner_power.TotalPower.RawBytePower
			QualityAdjPower = miner_power.TotalPower.QualityAdjPower

			balance, err := Filecoin.WalletBalance(data["miner"])
			if err != nil {
				Log.Crrs(err, tuuz.FUNCTION_ALL())
				continue
			} else {

			}
			baseinfo, err := Filecoin.MinerGetBaseInfo(data["miner"])
			if err != nil {
				Log.Crrs(err, tuuz.FUNCTION_ALL())
				continue
			} else {

			}
			MinerModel.Api_update_power(data["miner"],
				miner_power.MinerPower.RawBytePower,
				miner_power.MinerPower.QualityAdjPower,
				miner_power.HasMinPower, baseinfo.WorkerKey,
				baseinfo.SectorSize)

			MinerPowerLogModel.Api_insert(data["miner"],
				miner_power.MinerPower.RawBytePower,
				miner_power.MinerPower.QualityAdjPower,
				miner_power.HasMinPower, balance,
				baseinfo.WorkerKey, baseinfo.SectorSize,
				baseinfo.MinerPower, baseinfo.NetworkPower,
				time.Now().Unix())
		}
	}
	SystemParamModel.Api_update("RawBytePower", RawBytePower)
	SystemParamModel.Api_update("QualityAdjPower", QualityAdjPower)
}

func refresh_miner_amount() {
	datas := MinerModel.Api_select_all()
	var mm MinerModel.Interface
	mm.Db = tuuz.Db()
	mm.Db.Begin()
	for _, data := range datas {
		balance, err := Filecoin.WalletBalance(data["miner"])
		if err != nil {
			Log.Crrs(err, tuuz.FUNCTION_ALL())
		} else {
			//fmt.Println(data["miner"], balance)
			mm.Api_update_amount(data["miner"], balance)
		}
	}
	mm.Db.Commit()
}

func refresh_state() {
	state, err := Filecoin.StateVMCirculatingSupplyInternal()
	if err == nil {
		state.Timestamp = time.Now().Unix()
		SystemParamModel.Api_update("FilBurnt", state.FilBurnt)
		SystemParamModel.Api_update("FilLocked", state.FilLocked)
		SystemParamModel.Api_update("FilCirculating", state.FilCirculating)
		SystemParamModel.Api_update("FilMined", state.FilMined)
		SystemParamModel.Api_update("FilReserveDisbursed", state.FilReserveDisbursed)
		SystemParamModel.Api_update("FilVested", state.FilVested)
		CirculatingModel.Api_insert_manual(state)
	}
}
