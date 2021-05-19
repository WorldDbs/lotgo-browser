package MessageBlsModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "message_bls"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(belong_cid, height, miner, message_cid, Version, From, To, Nonce, Value, GasLimit, GasFeeCap, GasPremium, Method, Params, date interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"belong_cid":  belong_cid,
		"height":      height,
		"miner":       miner,
		"message_cid": message_cid,
		"Version":     Version,
		"From":        From,
		"To":          To,
		"Nonce":       Nonce,
		"Value":       Value,
		"GasLimit":    GasLimit,
		"GasFeeCap":   GasFeeCap,
		"GasPremium":  GasPremium,
		"Method":      Method,
		"Params":      Params,
		"date":        date,
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

func Api_select(belong_cid interface{}, page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	where := map[string]interface{}{
		"belong_cid": belong_cid,
	}
	db.Where(where)
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

func Api_select_all(page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
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

func Api_select_gas_all() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("Method,sum(GasLimit) as GasLimit,sum(GasFeeCap) as GasFeeCap,sum(GasPremium) as GasPremium,date as timestamp")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_gas_groupMethod() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("sum(GasLimit) as GasLimit,sum(GasFeeCap) as GasFeeCap,sum(GasPremium) as GasPremium,date as timestamp")
	db.Group("Method")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_count_all() int64 {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	db.Order("id desc")
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_select_from(From interface{}, page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	where := map[string]interface{}{
		"`From`": From,
	}
	db.Where(where)
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

func Api_select_to(To interface{}, page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	where := map[string]interface{}{
		"To": To,
	}
	db.Where(where)
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

func Api_select_byMethod(method []interface{}) interface{} {
	db := tuuz.Db().Table(table)
	db.Fields("SUM(Value) as val")
	db.WhereIn("Method", method)
	db.Order("id desc")
	ret, err := db.Value("val")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_last(address interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	db.OrWhere("From", address)
	db.OrWhere("to", address)
	db.Order("id desc")
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_early(address interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	db.OrWhere("From", address)
	db.OrWhere("to", address)
	db.OrWhere("miner", address)
	db.Order("id desc")
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_byAddress(address interface{}, page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	db.Where("From", address)
	db.OrWhere("To", address)
	db.OrWhere("miner", address)
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

func Api_count_byAddress(address interface{}) int64 {
	db := tuuz.Db().Table(table)
	db.Where("From", address)
	db.OrWhere("To", address)
	db.OrWhere("miner", address)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_count_byBelongCid(belong_cid interface{}) int64 {
	db := tuuz.Db().Table(table)
	db.OrWhere("belong_cid", belong_cid)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_sum_value(belong_cid interface{}) interface{} {
	db := tuuz.Db().Table(table)
	db.Where("belong_cid", belong_cid)
	db.Where("Method", 0)
	ret, err := db.Sum("Value")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_sum_destroy() interface{} {
	db := tuuz.Db().Table(table)
	ret, err := db.Sum("GasPremium")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_sum_To_byMethod(Method, To interface{}) interface{} {
	db := tuuz.Db().Table(table)
	db.Where("To", To)
	db.Where("Method", Method)
	ret, err := db.Sum("Value")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_sum_From_byMethod(Method, From interface{}) interface{} {
	db := tuuz.Db().Table(table)
	db.Where("From", From)
	db.Where("Method", Method)
	ret, err := db.Sum("Value")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_select_method(address, method interface{}, page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*,date as timestamp")
	db.OrWhere("From", address)
	db.OrWhere("To", address)
	db.OrWhere("miner", address)
	db.Where("Method", method)
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
