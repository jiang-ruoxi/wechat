package pkg

import (
	"encoding/json"
	"log"
	"wechat/config"
	"wechat/router"
	"wechat/utils"
)

func Execute() {
	//获取配置文件
	conf := config.InitConfig()
	//初始化路由
	r := router.InitRouter()
	//初始化DB
	InitDB(&conf)
	//获取客户端访问ip信息
	executeIPInfo()
	//启动WEB服务
	if err := r.Run(":" + conf.System.Port); err != nil {
		log.Fatal("服务器启动失败...")
	}
}

func executeIPInfo() {
	ipAddress, _ := utils.GetIpAddress()
	log.Printf("客户端IP:%#v \n", ipAddress)

	ipInfo, _ := utils.GetIPDataInfo(ipAddress)
	ipData, _ := json.Marshal(ipInfo)
	log.Printf("客户端IP详情:%#v \n", string(ipData))
}
