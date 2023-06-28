package crontab

import "sdwebuiapi/service"

// SdQueue sd 排队队列
func SdQueue() {
	sdService := service.SdService{}
	sdService.SyncQueue()
	sdService.SyncQueueClear()
}
