package config

type Wechat struct {
	AppId  string `mapstructure:"appid" json:"appid" yaml:"appid"`    // appid
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 端口值
}
