package model

// ShuXue 结构体
type ShuXue struct {
	Id         int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	AddTime    string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间;size:255;"`
}

// TableName BaiKe 表名
func (ShuXue) TableName() string {
	return "s_shuxue"
}
