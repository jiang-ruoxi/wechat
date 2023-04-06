package model

// Category 结构体
type Category struct {
	Id      int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:10;"`
	Name    string `json:"name" form:"name" gorm:"column:name;comment:栏目名称;size:255;"`
	//Status  string `json:"status" form:"status" gorm:"column:status;comment:状态,1启用,0禁用;size:255;"`
	AddTime string `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间;size:255;"`
}

// TableName Category 表名
func (Category) TableName() string {
	return "s_category"
}
