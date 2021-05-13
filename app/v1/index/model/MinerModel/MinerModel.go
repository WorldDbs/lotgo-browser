package MinerModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "miner"

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
