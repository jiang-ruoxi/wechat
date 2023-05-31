package model

type BookInfo struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   int    `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Mp3      string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:mp3;size:1024;"`
	Pic      string `json:"pic" form:"pic" gorm:"column:pic;comment:详情图;size:1024;"`
	En       string `json:"en" form:"en" gorm:"column:en;comment:英文内容;size:1024;"`
	Zh       string `json:"zh" form:"zh" gorm:"column:zh;comment:中文内容;size:1024;"`
	Duration uint8  `json:"duration" form:"duration" gorm:"column:duration;comment:时长;size:3;"`
	Position uint8  `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
}

// TableName Book 表名
func (BookInfo) TableName() string {
	return "s_huiben_info"
}
