package MinerModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "miner"

type Interface struct {
	Db gorose.IOrm
}

func Api_insert(miner interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"miner": miner,
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

func Api_delete(miner interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"miner": miner,
	}
	db.Where(where)
	_, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select(page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Limit(limit)
	db.Page(page)
	db.Order("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_all() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Order("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_sum_QualityAdjPower() interface{} {
	db := tuuz.Db().Table(table)
	ret, err := db.Sum("QualityAdjPower")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_sum_RawBytePower() interface{} {
	db := tuuz.Db().Table(table)
	ret, err := db.Sum("RawBytePower")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_count() int64 {
	db := tuuz.Db().Table(table)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_update_power(miner, RawBytePower, QualityAdjPower, HasMinPower, WorkerKey, SectorSize interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"miner": miner,
	}
	db.Where(where)
	data := map[string]interface{}{
		"RawBytePower":    RawBytePower,
		"QualityAdjPower": QualityAdjPower,
		"HasMinPower":     HasMinPower,
		"WorkerKey":       WorkerKey,
		"SectorSize":      SectorSize,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_update_amount(miner, amount interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"miner": miner,
	}
	db.Where(where)
	data := map[string]interface{}{
		"amount": amount,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
