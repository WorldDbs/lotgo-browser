package cron

import (
	"main.go/extend/Filecoin"
	"time"
)

func Refresh_gas() {
	tick := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tick.C:
			Filecoin.GasEstimateGasPremium()
			Filecoin.StateMinerInitialPledgeCollateral()
		}
	}
}
