package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Mysql struct {
	Path        string `json:"path" yaml:"path"`
	Port        string `json:"port" yaml:"port"`
	Config      string `json:"config" yaml:"config"`
	Db          string `json:"db" yaml:"db"`
	UserName    string `json:"user_name" yaml:"user_name"`
	Password    string `json:"password" yaml:"password"`
	MaxIdleConn int `json:"max_idle_conn" yaml:"max_idle_conn"`
	MaxOpenConn int `json:"max_open_conn" yaml:"max_open_conn"`
}

type Redis struct {
	DB       int `json:"db" yaml:"db"`
	Addr     string `json:"addr" yaml:"addr"`
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
}

type System struct {
	Port string `json:"port" yaml:"port"`
}

type Config struct {
	System System `json:"system"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	Redis  Redis  `json:"redis" yaml:"redis"`
}

func InitConfig() (conf Config) {
	path := "./config.yaml"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("配置文件加载失败")
		return
	}
	//yaml.Unmarshal会根据yaml标签的字段进行赋值
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal("配置文件Unmarshal failed")
		return
	}
	return
}
