package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/MessageBlsModel"
	"main.go/app/v1/wallet/model/WalletUserModel"
	"main.go/common/BaseController"
	"main.go/extend/Filecoin"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func TransferController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("send", transfer_send)
	route.Any("recv", transfer_receive)
	route.Any("list_recv", transfer_list_recv)
	route.Any("list_send", transfer_list_send)
	//route.Any("list", transfer_list)
}

func transfer_send(c *gin.Context) {
	token := c.PostForm("token")
	user := WalletUserModel.Api_find_byToken(token)
	//fmt.Println(user)
	to, ok := Input.Post("to", c, false)
	if !ok {
		return
	}
	value, ok := Input.Post("value", c, false)
	if !ok {
		return
	}
	gaslimit, ok := Input.PostInt64("gaslimit", c)
	if !ok {
		return
	}
	data, err := Filecoin.MpoolPushMessage(user["address"], to, value, gaslimit)
	if err != nil {
		RET.Fail(c, 300, err.Error(), err.Error())
	} else {
		RET.Success(c, 0, data, nil)
	}
}

func transfer_receive(c *gin.Context) {
	token := c.PostForm("token")
	data := WalletUserModel.Api_find_byToken(token)
	RET.Success(c, 0, data["address"], nil)
}

func transfer_list_recv(c *gin.Context) {
	token := c.PostForm("token")
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit, ok := Input.PostInt("limit", c)
	if !ok {
		return
	}
	wallet := WalletUserModel.Api_find_byToken(token)
	data := MessageBlsModel.Api_select_to(wallet["address"], page, limit)
	RET.Success(c, 0, data, nil)
}

func transfer_list_send(c *gin.Context) {
	token := c.PostForm("token")
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit, ok := Input.PostInt("limit", c)
	if !ok {
		return
	}
	wallet := WalletUserModel.Api_find_byToken(token)
	data := MessageBlsModel.Api_select_from(wallet["address"], page, limit)
	RET.Success(c, 0, data, nil)
}
