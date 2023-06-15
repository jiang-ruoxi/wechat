package model

type ChineseBookInfo struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   int    `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Mp3      string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:mp3;size:1024;"`
	Pic      string `json:"pic" form:"pic" gorm:"column:pic;comment:详情图;size:1024;"`
	Position uint8  `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
}

// TableName ChineseBookInfo 表名
func (ChineseBookInfo) TableName() string {
	return "s_chinese_picture_info"
}
