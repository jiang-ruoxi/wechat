package model

type ChineseBook struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   string `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:绘本id;size:1024;"`
	Icon     string `json:"icon" form:"icon" gorm:"column:icon;comment:封面图;size:1024;"`
	Level    uint8  `json:"level" form:"level" gorm:"column:level;comment:级别;size:1;"`
	Position uint8  `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
}

// TableName Book 表名
func (ChineseBook) TableName() string {
	return "s_chinese_picture"
}
