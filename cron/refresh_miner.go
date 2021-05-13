package cron

import (
	"main.go/extend/Filecoin"
	"time"
)

func Refresh_miner() {
	tick := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tick.C:
			refresh_miner()
		}
	}
}

func refresh_miner() {
	Filecoin.StateListMiners()
}
