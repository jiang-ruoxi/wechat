package pkg

import (
	"wechat/config"
	"wechat/pkg/mysql"
	"wechat/pkg/redis"
)

func InitDB(conf *config.Config) {
	//初始化连接Mysql
	mysql.InitDB(config.Config{
		Mysql: config.Mysql{
			Path:        conf.Mysql.Path,
			Port:        conf.Mysql.Port,
			Config:      conf.Mysql.Config,
			Db:          conf.Mysql.Db,
			UserName:    conf.Mysql.UserName,
			Password:    conf.Mysql.Password,
			MaxIdleConn: conf.Mysql.MaxIdleConn,
			MaxOpenConn: conf.Mysql.MaxOpenConn,
		},
	})
	//初始化连接Redis
	redis.InitDB(config.Config{
		Redis: config.Redis{
			DB:       conf.Redis.DB,
			Port:     conf.Redis.Port,
			Addr:     conf.Redis.Addr,
			Password: conf.Redis.Password,
		},
	})
}
