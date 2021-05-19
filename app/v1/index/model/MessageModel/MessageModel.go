package MessageModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "message"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(cid, height, datas, num_bls, num_secpk, num_cid interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"cid":       cid,
		"height":    height,
		"datas":     datas,
		"num_bls":   num_bls,
		"num_secpk": num_secpk,
		"num_cid":   num_cid,
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

func Api_find(cid interface{}) gorose.Data {
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

func Api_select(page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Order("id desc")
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_height(height interface{}) gorose.Data {
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

func Api_count_24() interface{} {
	db := tuuz.Db().Table(table)
	ret, err := db.Query("SELECT SUM(num_cid) as num FROM `message` order by id desc LIMIT 17280")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret[0]["num"]
	}
}
