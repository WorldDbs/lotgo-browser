package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/GasModel"
	"main.go/tuuz/RET"
	"time"
)

func GasController(route *gin.RouterGroup) {
	route.Any("index", gas_index)
	route.Any("24", gas_24)
	route.Any("month", gas_month)
}

func gas_index(c *gin.Context) {

}

func gas_24(c *gin.Context) {
	time_start := time.Now().Unix() - 86400
	time_end := time.Now().Unix()
	data := GasModel.Api_manual_time(1800, time_start, time_end)
	RET.Success(c, 0, data, nil)
}

func gas_month(c *gin.Context) {
	time_start := time.Now().Unix() - 86400*30
	time_end := time.Now().Unix()
	data := GasModel.Api_manual_time(86400, time_start, time_end)
	RET.Success(c, 0, data, nil)
}
