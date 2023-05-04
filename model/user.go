package model

// User 结构体
type User struct {
	Id         int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:10;"`
	OpenId     string `json:"open_id" form:"open_id" gorm:"column:open_id;comment:用户open_id;size:10;"`
	NickName   string `json:"nick_name" form:"nick_name" gorm:"column:nick_name;comment:用户昵称;size:255;"`
	HeadUrl    string `json:"head_url" form:"head_url" gorm:"column:head_url;comment:用户头像;size:255;"`
	Area       string `json:"area" form:"area" gorm:"column:area;comment:地区;size:255;"`
	Score      int    `json:"score" form:"score" gorm:"column:score;comment:得分;size:11;"`
	AddTime    string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间;size:255;"`
	UpdateTime string `json:"update_time" form:"update_time" gorm:"column:update_time;comment:更新时间;size:255;"`
}

// TableName User 表名
func (User) TableName() string {
	return "s_user"
}
