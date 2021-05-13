package cron

import (
	"main.go/app/v1/index/model/HeightModel"
	"main.go/extend/Filecoin"
)

func NewBlock() {
	for {
		height := HeightModel.Api_find_last()
		if height == nil {
			height = int64(0)
		} else {
			height = height.(int64) + 1
		}
		Filecoin.Get_block_by_height(height.(int64))
	}
}
