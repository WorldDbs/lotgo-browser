package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/index/model/GasModel"
	"main.go/tuuz/RET"
	"time"
)

func ChartController(route *gin.RouterGroup) {
	route.Any("24", chart_24h)
	route.Any("7d", chart_7d)
	route.Any("month", chart_month)
	route.Any("year", chart_year)

}

func chart_24h(c *gin.Context) {
	time_start := time.Now().Unix() - 86400
	time_end := time.Now().Unix()
	data := GasModel.Api_manual_time(1800, time_start, time_end)
	RET.Success(c, 0, data, nil)
}

func chart_7d(c *gin.Context) {
	time_start := time.Now().Unix() - 86400*30
	time_end := time.Now().Unix()
	data := GasModel.Api_manual_time(43200, time_start, time_end)
	RET.Success(c, 0, data, nil)
}

func chart_year(c *gin.Context) {
	time_start := time.Now().Unix() - 86400*30
	time_end := time.Now().Unix()
	data := GasModel.Api_manual_time(86400*30, time_start, time_end)
	RET.Success(c, 0, data, nil)
}

func chart_month(c *gin.Context) {
	time_start := time.Now().Unix() - 86400*30
	time_end := time.Now().Unix()
	data := GasModel.Api_manual_time(86400, time_start, time_end)
	RET.Success(c, 0, data, nil)
}
