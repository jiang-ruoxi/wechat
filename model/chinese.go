package model

// ChineseBook 中文国学绘本对应的书籍
type ChineseBook struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Type     uint8  `json:"type"`
	Position uint8  `json:"position"`
}

// TableName ChineseBook 表名
func (ChineseBook) TableName() string {
	return "s_chinese_picture"
}

// ChineseBook 中文国学绘本对应的书籍具体的详情
type ChineseBookInfo struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	Position uint8  `json:"position"`
}

// TableName ChineseBookInfo 表名
func (ChineseBookInfo) TableName() string {
	return "s_chinese_picture_info"
}

//PoetryPicture 古诗绘本
type PoetryPicture struct {
	Id     int    `json:"-"`
	BookId int    `json:"book_id"`
	Title  string `json:"title"`
	Icon   string `json:"icon"`
	Author string `json:"author"`
	TypeId int    `json:"type_id"`
}

// TableName PoetryPicture 表名
func (PoetryPicture) TableName() string {
	return "s_poetry_picture"
}

//PoetryPictureInfo 古诗绘本详情
type PoetryPictureInfo struct {
	Id       int    `json:"-"`
	BookId   int    `json:"book_id"`
	CN       string `json:"cn"`
	Pic      string `json:"pic"`
	Mp3      string `json:"mp3"`
	Position int    `json:"position"`
}

// TableName PoetryPictureInfo 表名
func (PoetryPictureInfo) TableName() string {
	return "s_poetry_picture_info"
}
