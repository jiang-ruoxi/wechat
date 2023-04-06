package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"wechat/config"
)

var DB *gorm.DB

func InitDB(conf config.Config) {
	var err error
	dbParams := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",
		conf.Mysql.UserName,
		conf.Mysql.Password,
		conf.Mysql.Path,
		conf.Mysql.Port,
		conf.Mysql.Db,
		conf.Mysql.Config,
	)
	DB, err = gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatal("mysql数据库连接失败：", err)
	}
	// 全局禁用表名复数
	DB.SingularTable(true)

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	fmt.Println("database init on port ", conf.Mysql.Port)
}
