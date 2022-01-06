package main

import (
	"beamwang/day06/mylogger"
	"time"
)

var log mylogger.Logger //声明一个全局的接口变量

// 测试我们自己写的日志库
func main() {
	// log := mylogger.NewLog("debug")
	id := 1510260111
	name := "beam"
	log = mylogger.NewConsoleLog("debug")
	log = mylogger.NewFileLog("debug", "./log", 2*1024)
	for {
		log.Debug("这是一条Debug日志: id:%d, name:%s", id, name)
		log.Info("这是一条Info日志")
		log.Error("这是一条Error日志")
		log.Warning("这是一条Warning日志")
		time.Sleep(time.Second * 2)
	}
}
