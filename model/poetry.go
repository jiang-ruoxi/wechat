package model

type Poetry struct {
	Id         int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	PoetryId   int    `json:"poetry_id" form:"poetry_id" gorm:"column:poetry_id;comment:古诗id;size:10;"`
	Title      string `json:"title" form:"title" gorm:"column:title;comment:标题;size:256;"`
	GradeId    uint8  `json:"grade_id" form:"grade_id" gorm:"column:grade_id;comment:年级id;size:1;"`
	Grade      string `json:"grade" form:"grade" gorm:"column:grade;comment:年级;size:64;"`
	GradeLevel uint8  `json:"grade_level" form:"grade_level" gorm:"column:grade_level;comment:年级上下;size:1;"`
	Author     string `json:"author" form:"author" gorm:"column:author;comment:作者;size:64;"`
	Dynasty    string `json:"dynasty" form:"dynasty" gorm:"column:dynasty;comment:朝代;size:64;"`
	Mp3        string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:古诗音频;size:256;"`
	Content    string `json:"content" form:"content" gorm:"column:content;comment:古诗内容;size:1024;"`
	Info       string `json:"info" form:"info" gorm:"column:info;comment:古诗拼音;size:1024;"`
}

// TableName Poetry 表名
func (Poetry) TableName() string {
	return "s_poetry"
}
