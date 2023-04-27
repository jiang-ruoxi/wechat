package model

// Answer 结构体
type Answer struct {
	Id          int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:10;"`
	CategoryId  string `json:"category_id" form:"category_id" gorm:"column:category_id;comment:栏目Id;size:10;"`
	QuestionId  string `json:"question_id" form:"question_id" gorm:"column:question_id;comment:问题Id;size:255;"`
	OpenId      string `json:"open_id" form:"open_id" gorm:"column:open_id;comment:用户OpenId;size:10;"`
	IsSelect    string `json:"is_select" form:"is_select" gorm:"column:is_select;comment:选择的答案;size:255;"`
	RightSelect string `json:"right_select" form:"right_select" gorm:"column:right_select;comment:正确选项;size:255;"`
	AddTime     string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间;size:255;"`
}

// TableName Answer 表名
func (Answer) TableName() string {
	return "s_answer_log"
}
