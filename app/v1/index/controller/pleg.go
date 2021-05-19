package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/PlegModel"
	"main.go/tuuz/RET"
	"time"
)

func PlegController(route *gin.RouterGroup) {
	route.Any("index", pleg_index)
	route.Any("24", pleg_24)
	route.Any("month", pleg_month)
}

func pleg_index(c *gin.Context) {

}

func pleg_24(c *gin.Context) {
	time_start := time.Now().Unix() - 86400
	time_end := time.Now().Unix()
	data := PlegModel.Api_manual_time(1800, time_start, time_end)
	RET.Success(c, 0, data, nil)
}

func pleg_month(c *gin.Context) {
	time_start := time.Now().Unix() - 86400*30
	time_end := time.Now().Unix()
	data := PlegModel.Api_manual_time(86400, time_start, time_end)
	RET.Success(c, 0, data, nil)
}
