package MessageSecpkModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "message_speck"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(belong_cid, height, message_cid, Version, From, To, Nonce, Value, GasLimit, GasFeeCap, GasPremium, Method, Params, Signature, CID interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"belong_cid":  belong_cid,
		"height":      height,
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
		"Signature":   Signature,
		"CID":         CID,
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
