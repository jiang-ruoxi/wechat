package main

import (
	"go.uber.org/zap"
	"wechat/pkg"
	"wechat/pkg/log"
)

func init() {
	// 初始化日志库
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, "./log/gin-example.log")
}

func main() {
	pkg.Execute()
}
