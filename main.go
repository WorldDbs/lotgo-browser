package main

import (
	"github.com/gin-gonic/gin"
	"main.go/cron"
	"main.go/route"
)

func main() {

	go cron.NewBlock()
	go cron.Refresh_miner()
	go cron.Refresh_gas()

	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
