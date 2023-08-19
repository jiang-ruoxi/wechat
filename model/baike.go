package model

// Category 结构体
type Category struct {
	Id   int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:10;"`
	Name string `json:"name" form:"name" gorm:"column:name;comment:栏目名称;size:255;"`
}

// TableName Category 表名
func (Category) TableName() string {
	return "s_category"
}

// BaiKe 结构体
type BaiKe struct {
	Id         int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:10;"`
	CategoryId int    `json:"category_id" form:"category_id" gorm:"column:category_id;comment:栏目Id;size:10;"`
	Question   string `json:"question" form:"question" gorm:"column:question;comment:问题;size:255;"`
	OptionA    string `json:"option_a" form:"option_a" gorm:"column:option_a;comment:选项a;size:255;"`
	OptionB    string `json:"option_b" form:"option_b" gorm:"column:option_b;comment:选项b;size:255;"`
	OptionC    string `json:"option_c" form:"option_c" gorm:"column:option_c;comment:选项c;size:255;"`
	OptionD    string `json:"option_d" form:"option_d" gorm:"column:option_d;comment:选项d;size:255;"`
	Answer     string `json:"answer" form:"answer" gorm:"column:answer;comment:答案;size:255;"`
	Analytic   string `json:"analytic" form:"analytic" gorm:"column:analytic;comment:原因;size:255;"`
	AddTime    string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间;size:255;"`
}

// TableName BaiKe 表名
func (BaiKe) TableName() string {
	return "s_baike"
}
