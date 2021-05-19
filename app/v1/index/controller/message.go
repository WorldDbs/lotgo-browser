package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/MessageBlsModel"
	"main.go/app/v1/index/model/MessageSecpkModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func MessageController(route *gin.RouterGroup) {
	route.Any("list", message_list)
	route.Any("bls_list", message_bls_list)
	route.Any("secpk_list", message_secpk_list)
}

func message_list(c *gin.Context) {
	page, ok := Input.PostInt("page", c)
	if !ok {
		return
	}
	limit, ok := Input.PostInt("limit", c)
	if !ok {
		return
	}
	datas := MessageBlsModel.Api_select_all(page, limit)
	for i, data := range datas {
		data["timestamp"] = data["date"]
		datas[i] = data
	}
	RET.Success(c, 0, datas, MessageBlsModel.Api_count_all())
}

func message_bls_list(c *gin.Context) {
	cid, ok := Input.Post("cid", c, false)
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
	count := MessageBlsModel.Api_count_byBelongCid(cid)
	datas := MessageBlsModel.Api_select(cid, page, limit)
	RET.Success(c, 0, datas, count)
}

func message_secpk_list(c *gin.Context) {
	cid, ok := Input.Post("cid", c, false)
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
	datas := MessageSecpkModel.Api_select(cid, page, limit)
	for i, data := range datas {
		data["timestamp"] = data["date"]
		datas[i] = data
	}
	RET.Success(c, 0, datas, nil)
}
