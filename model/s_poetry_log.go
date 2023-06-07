package model

// PoetryLog 结构体
type PoetryLog struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	OpenId   string `json:"open_id" form:"open_id" gorm:"column:open_id;comment:用户open_id;size:128;"`
	PoetryId   int    `json:"poetry_id" form:"poetry_id" gorm:"column:poetry_id;comment:古诗id;size:11;"`
	Mp3      string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:转换后的url;size:256;"`
}

func (PoetryLog) TableName() string {
	return "s_poetry_log"
}
