package cron

import (
	"runtime"

	"github.com/jasonlvhit/gocron"
	"github.com/8812488/proxypool/internal/app"
)

func Cron() {
	_ = gocron.Every(15).Minutes().Do(crawlTask)
	<-gocron.Start()
}

func crawlTask() {
	_ = app.InitConfigAndGetters("")
	app.CrawlGo()
	app.Getters = nil
	runtime.GC()
}
