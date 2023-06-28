package crontab

import (
	"github.com/robfig/cron"
)

func Init() {
	c := cron.New()

	c.AddFunc("* * * * *", func() {
		//  SdQueue 排队队列
		SdQueue()
	})

	c.Start()

}
