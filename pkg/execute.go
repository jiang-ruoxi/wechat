package pkg

import (
	"log"
	"wechat/config"
	"wechat/router"
)

func Execute() {
	//获取配置文件
	conf := config.InitConfig()
	//初始化路由
	r := router.InitRouter()
	//初始化DB
	InitDB(&conf)
	//启动WEB服务
	if err := r.Run(":" + conf.System.Port); err != nil {
		log.Fatal("服务器启动失败...")
	}
}
