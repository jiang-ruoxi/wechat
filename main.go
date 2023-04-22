package main

import (
	"go.uber.org/zap"
	"wechat/pkg"
	"wechat/pkg/log"
	"wechat/utils"
)

func init() {
	// 初始化日志库
	file := utils.GetCurrentFormatDateTime() + ".log"
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, log.LOG_DIR+file)
}

func main() {
	pkg.Execute()
}
