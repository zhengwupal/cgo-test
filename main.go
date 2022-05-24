package main

import (
	"time"

	logs "github.com/zhengwupal/cgo-test/logs"
	pack1 "github.com/zhengwupal/cgo-test/pack1"
)

func main() {
	log_level := "info"
	logs.InitLog(log_level+".log", "error.log", log_level)
	defer logs.ZapLogger.Sync()
	logs.SugarLogger.Infof("Start Run...")
	startTime := time.Now()

	// pack1.Add()
	// pack1.Test2()
	// pack1.Test3()
	// pack1.Test4()
	pack1.Test5()
	// pack1.Test6()
	// pack1.Test7()

	cost := time.Since(startTime) / 1000000
	logs.SugarLogger.Infof("Success! Consume time %dms", cost)
}
