package model

// EnglishBook 英语绘本对应的书籍
type EnglishBook struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   string `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:绘本id;size:1024;"`
	Icon     string `json:"icon" form:"icon" gorm:"column:icon;comment:封面图;size:1024;"`
	Level    uint8  `json:"level" form:"level" gorm:"column:level;comment:级别;size:1;"`
	Position uint8  `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
}

// TableName EnglishBook 表名
func (EnglishBook) TableName() string {
	return "s_english_picture"
}

// EnglishBookInfo 英语绘本对应的书籍具体的详情
type EnglishBookInfo struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   string `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Mp3      string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:mp3;size:1024;"`
	Pic      string `json:"pic" form:"pic" gorm:"column:pic;comment:详情图;size:1024;"`
	En       string `json:"en" form:"en" gorm:"column:en;comment:英文内容;size:1024;"`
	Zh       string `json:"zh" form:"zh" gorm:"column:zh;comment:中文内容;size:1024;"`
	Duration uint8  `json:"duration" form:"duration" gorm:"column:duration;comment:时长;size:3;"`
	Position uint8  `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
}

// TableName EnglishBookInfo 表名
func (EnglishBookInfo) TableName() string {
	return "s_english_picture_info"
}
