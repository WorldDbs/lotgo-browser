package WalletUserModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "wallet_user"

func Api_insert(address, private_key, password, token interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"address":     address,
		"private_key": private_key,
		"password":    password,
		"token":       token,
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

func Api_find(address interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"address": address,
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

func Api_find_byPassword(address, password interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"address":  address,
		"password": password,
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

func Api_find_byAddressAndToken(address, token interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"address": address,
		"token":   token,
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

func Api_find_byToken(token interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"token": token,
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

func Api_update_byAddress(address, password, token interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"address": address,
	}
	db.Where(where)
	data := map[string]interface{}{
		"password": password,
		"token":    token,
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

func Api_update_token(address, new_token interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"address": address,
	}
	db.Where(where)
	data := map[string]interface{}{
		"token": new_token,
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

func Api_update_byToken(token, new_token interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"token": token,
	}
	db.Where(where)
	data := map[string]interface{}{
		"token": new_token,
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
