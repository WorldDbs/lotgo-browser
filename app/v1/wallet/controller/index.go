package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/BlocksModel"
	"main.go/app/v1/index/model/MessageBlsModel"
	"main.go/app/v1/wallet/model/WalletUserModel"
	"main.go/extend/Filecoin"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("new", index_new)
	route.Any("relogin", index_relogin)
	route.Any("login", index_login)
	route.Any("balance", index_balance)
	route.Any("info", index_info)
	route.Any("user_message", index_user_message)
	route.Any("user_block", index_user_block)
	route.Any("user_transfer", index_user_transfer)
}

func index_new(c *gin.Context) {
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}

	bls, pvk, token, err := Filecoin.WalletNew(password)
	if err != nil {
		RET.Fail(c, 300, err.Error(), err.Error())
	} else {
		RET.Success(c, 0, map[string]interface{}{
			"bls":         bls,
			"private_key": pvk,
			"token":       token,
		}, nil)
	}
}

func index_relogin(c *gin.Context) {
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	if len(WalletUserModel.Api_find_byPassword(address, password)) > 0 {
		new_token := Calc.GenerateToken()
		if WalletUserModel.Api_update_token(address, new_token) {
			RET.Success(c, 0, new_token, nil)
		} else {
			RET.Fail(c, 500, nil, nil)
		}
	} else {
		RET.Fail(c, -1, nil, nil)
	}

}

func index_login(c *gin.Context) {
	private_key, ok := Input.Post("private_key", c, false)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	address, err := Filecoin.WalletImport(private_key)
	if err != nil {
		RET.Fail(c, 300, err.Error(), err.Error())
		return
	}
	token := Calc.GenerateToken()
	user := WalletUserModel.Api_find(address)
	if len(user) > 0 {
		if WalletUserModel.Api_update_byAddress(address, password, token) {
			RET.Success(c, 0, map[string]interface{}{
				"bls":         address,
				"private_key": private_key,
				"token":       token,
			}, nil)
			return
		}
	} else {
		if WalletUserModel.Api_insert(address, private_key, password, token) {
			RET.Success(c, 0, map[string]interface{}{
				"bls":         address,
				"private_key": private_key,
				"token":       token,
			}, nil)
			return
		}
	}
	RET.Fail(c, 500, nil, nil)
}

func index_balance(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	data, err := Filecoin.WalletBalance(address)
	if err != nil {
		RET.Fail(c, 300, data, nil)
	} else {
		RET.Success(c, 0, data, nil)
	}
}

func index_info(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	data, err := Filecoin.WalletBalance(address)
	last := MessageBlsModel.Api_find_last(address)
	early := MessageBlsModel.Api_find_early(address)
	if err != nil {
		RET.Fail(c, 300, nil, nil)
	} else {
		RET.Success(c, 0, map[string]interface{}{
			"start_time": early["timestamp"],
			"last_time":  last["timestamp"],
			"balance":    data,
		}, nil)
	}
}

func index_user_transfer(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit, ok := Input.PostInt("limit", c)
	if !ok {
		return
	}
	datas := MessageBlsModel.Api_select_byAddress(address, page, limit)
	count := MessageBlsModel.Api_count_byAddress(address)
	RET.Success(c, 0, datas, count)
}

func index_user_block(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit, ok := Input.PostInt("limit", c)
	if !ok {
		return
	}
	datas := MessageBlsModel.Api_select_byAddress(address, page, limit)
	count := MessageBlsModel.Api_count_byAddress(address)
	arr := []interface{}{}
	for _, data := range datas {
		temp := BlocksModel.Api_select_byHeight(data["height"])
		arr = append(arr, temp)
	}
	RET.Success(c, 0, arr, count)
}

func index_user_message(c *gin.Context) {
	address, ok := Input.Post("address", c, false)
	if !ok {
		return
	}
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit, ok := Input.PostInt("limit", c)
	if !ok {
		return
	}
	datas := MessageBlsModel.Api_select_byAddress(address, page, limit)
	count := MessageBlsModel.Api_count_byAddress(address)
	RET.Success(c, 0, datas, count)
}
