package BlocksModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Date"
	"main.go/tuuz/Log"
)

const table = "blocks"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(cid, height, timestamp, miner, parents, parent_weight, message, messages, datas interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"cid":           cid,
		"height":        height,
		"timestamp":     timestamp,
		"miner":         miner,
		"parents":       parents,
		"parent_weight": parent_weight,
		"message":       message,
		"messages":      messages,
		"datas":         datas,
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

func Api_find(height interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"height": height,
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

func Api_find_new() gorose.Data {
	db := tuuz.Db().Table(table)
	db.Order("id desc")
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
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

func Api_select_byHeight(height interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"height": height,
	}
	db.Where(where)
	ret, err := db.Get()
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

func Api_count_byMiner(miner interface{}) int64 {
	db := tuuz.Db().Table(table)
	db.Where("miner", miner)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_find_byCid(cid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"cid": cid,
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

func Api_count_24() int64 {
	db := tuuz.Db().Table(table)
	db.Where("timestamp", ">", Date.Today_int())
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}
