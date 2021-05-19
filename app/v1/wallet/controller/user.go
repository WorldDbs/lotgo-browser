package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/wallet/model/WalletUserModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func UserController(route *gin.RouterGroup) {

	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("export", user_export)
}

func user_export(c *gin.Context) {
	token := c.PostForm("token")
	data := WalletUserModel.Api_find_byToken(token)
	if len(data) > 0 {
		RET.Success(c, 0, data["private_key"], nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
