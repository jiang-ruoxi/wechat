package model

// ChineseBook 中文国学绘本对应的书籍
type ChineseBook struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   string `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:绘本id;size:1024;"`
	Icon     string `json:"icon" form:"icon" gorm:"column:icon;comment:封面图;size:1024;"`
	Type     uint8  `json:"type" form:"level" gorm:"column:type;comment:级别;size:1;"`
	Position uint8  `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
}

// TableName ChineseBook 表名
func (ChineseBook) TableName() string {
	return "s_chinese_picture"
}

// ChineseBook 中文国学绘本对应的书籍具体的详情
type ChineseBookInfo struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   string `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Mp3      string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:mp3;size:1024;"`
	Pic      string `json:"pic" form:"pic" gorm:"column:pic;comment:详情图;size:1024;"`
	Position uint8  `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
}

// TableName ChineseBookInfo 表名
func (ChineseBookInfo) TableName() string {
	return "s_chinese_picture_info"
}

//PoetryPicture 古诗绘本
type PoetryPicture struct {
	Id     int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId int    `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Title  string `json:"title" form:"title" gorm:"column:title;comment:title;size:1024;"`
	Icon   string `json:"icon" form:"icon" gorm:"column:icon;comment:缩略图;size:1024;"`
	Author string `json:"author" form:"author" gorm:"column:author;comment:作者;size:1024;"`
	TypeId int    `json:"type_id" form:"type_id" gorm:"column:type_id;comment:类型;size:3;"`
}

// TableName PoetryPicture 表名
func (PoetryPicture) TableName() string {
	return "s_poetry_picture"
}

//PoetryPictureInfo 古诗绘本详情
type PoetryPictureInfo struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	BookId   int    `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	CN       string `json:"cn" form:"cn" gorm:"column:cn;comment:cn;size:1024;"`
	Pic      string `json:"pic" form:"pic" gorm:"column:pic;comment:详情图;size:1024;"`
	Mp3      string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:mp3;size:1024;"`
	Position int    `json:"position" form:"position" gorm:"column:position;comment:排序;size:3;"`
}

// TableName PoetryPictureInfo 表名
func (PoetryPictureInfo) TableName() string {
	return "s_poetry_picture_info"
}
