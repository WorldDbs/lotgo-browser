package MinerPowerLogModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "miner_power_log"

func Api_insert(miner, RawBytePower, QualityAdjPower, HasMinPower, amount, WorkerKey, SectorSize, MinerPower, NetworkPower, timestamp interface {
}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"miner":           miner,
		"RawBytePower":    RawBytePower,
		"QualityAdjPower": QualityAdjPower,
		"HasMinPower":     HasMinPower,
		"amount":          amount,
		"WorkerKey":       WorkerKey,
		"SectorSize":      SectorSize,
		"MinerPower":      MinerPower,
		"NetworkPower":    NetworkPower,
		"timestamp":       timestamp,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(miner interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"miner": miner,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select(miner, start_timestamp, end_timestamp interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("miner", miner)
	db.Where("timestamp", "<", start_timestamp)
	db.Where("timestamp", ">", end_timestamp)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_one(miner interface{}, method, mode_desc_asc string, time interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("miner", miner)
	db.Where("timestamp", method, time)
	db.Order("id " + mode_desc_asc)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_sum_amount_all(time interface{}) interface{} {
	ret, err := tuuz.Db().Query("select sum(amount) as amount from (SELECT * FROM `miner_power_log` WHERE `timestamp` > ? GROUP BY miner) as a", time)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret[0]["amount"]
	}
}

func Api_manual_time(lap, time_start, time_end interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	ret, err := db.Query("SELECT *,floor(time/?) as timer FROM `gas` where time<? and time>? group by timer", lap, time_end, time_start)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
