package SystemParamModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "system_param"

func Api_find(key interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"key": key,
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

func Api_value(key interface{}) interface{} {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"key": key,
	}
	db.Where(where)
	ret, err := db.Value("value")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_update(key, value interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"key": key,
	}
	db.Where(where)
	data := map[string]interface{}{
		"value": value,
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
