package GasModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "gas"

func Api_insert(price, time interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"price": price,
		"time":  time,
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

func Api_select(time_start interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("time", ">", time_start)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
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
