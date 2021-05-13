package SystemParamModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "system_param"

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
