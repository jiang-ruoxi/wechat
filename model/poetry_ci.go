package model

type PoetryCI struct {
	Id         int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	PoetryId   int    `json:"poetry_id" form:"poetry_id" gorm:"column:poetry_id;comment:古诗id;size:10;"`
	Title      string `json:"title" form:"title" gorm:"column:title;comment:标题;size:256;"`
	GradeId    uint8  `json:"grade_id" form:"grade_id" gorm:"column:grade_id;comment:年级id;size:1;"`
	Grade      string `json:"grade" form:"grade" gorm:"column:grade;comment:年级;size:64;"`
	Author     string `json:"author" form:"author" gorm:"column:author;comment:作者;size:64;"`
	Dynasty    string `json:"dynasty" form:"dynasty" gorm:"column:dynasty;comment:朝代;size:64;"`
	Content    string `json:"content" form:"content" gorm:"column:content;comment:古诗内容;size:1024;"`
}

// TableName PoetryCI 表名
func (PoetryCI) TableName() string {
	return "s_junior_poetry"
}
