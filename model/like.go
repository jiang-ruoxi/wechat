package model

// Like 结构体
type Like struct {
	Id          int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:10;"`
	CategoryId  int    `json:"category_id" form:"category_id" gorm:"column:category_id;comment:栏目Id;size:10;"`
	QuestionId  int    `json:"question_id" form:"question_id" gorm:"column:question_id;comment:问题Id;size:255;"`
	OpenId      string `json:"open_id" form:"open_id" gorm:"column:open_id;comment:用户OpenId;size:10;"`
	AddTime     string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间;size:255;"`
}

// TableName Like 表名
func (Like) TableName() string {
	return "s_like"
}
