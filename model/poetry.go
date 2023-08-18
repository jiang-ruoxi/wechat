package model

// Poetry 小学级别的古诗词对应的表
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

// JuniorPoetry 初高中级别的古诗词对应的表
type JuniorPoetry struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	PoetryId int    `json:"poetry_id" form:"poetry_id" gorm:"column:poetry_id;comment:古诗id;size:10;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:标题;size:256;"`
	GradeId  uint8  `json:"grade_id" form:"grade_id" gorm:"column:grade_id;comment:年级id;size:1;"`
	Grade    string `json:"grade" form:"grade" gorm:"column:grade;comment:年级;size:64;"`
	Author   string `json:"author" form:"author" gorm:"column:author;comment:作者;size:64;"`
	Dynasty  string `json:"dynasty" form:"dynasty" gorm:"column:dynasty;comment:朝代;size:64;"`
	Content  string `json:"content" form:"content" gorm:"column:content;comment:古诗内容;size:1024;"`
}

// TableName JuniorPoetry 表名
func (JuniorPoetry) TableName() string {
	return "s_junior_poetry"
}

// ChengYU 成语对应的表
type ChengYU struct {
	Id      int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	Title   string `json:"title" form:"title" gorm:"column:title;comment:标题;size:1024;"`
	Pinyin  string `json:"pinyin" form:"pinyin" gorm:"column:pinyin;comment:拼音;size:1024;"`
	Explain string `json:"explain" form:"explain" gorm:"column:explain;comment:解释;size:1024;"`
	Source  string `json:"source" form:"source" gorm:"column:source;comment:出处;size:1024;"`
	Usage   string `json:"usage" form:"usage" gorm:"column:usage;comment:用法;size:1024;"`
	Example string `json:"example" form:"example" gorm:"column:example;comment:示例;size:1024;"`
	Near    string `json:"near" form:"near" gorm:"column:near;comment:近义词;size:1024;"`
	Antonym string `json:"antonym" form:"antonym" gorm:"column:antonym;comment:反义词;size:1024;"`
	Analyse string `json:"analyse" form:"analyse" gorm:"column:analyse;comment:辨析;size:1024;"`
	Story   string `json:"story" form:"story" gorm:"column:story;comment:故事;size:1024;"`
	Level   uint8  `json:"level" form:"level" gorm:"column:level;comment:级别;size:1;"`
}

// TableName ChengYU 表名
func (ChengYU) TableName() string {
	return "s_chengyu"
}

// PoetryLog 结构体
type PoetryLog struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	OpenId   string `json:"open_id" form:"open_id" gorm:"column:open_id;comment:用户open_id;size:128;"`
	PoetryId   int    `json:"poetry_id" form:"poetry_id" gorm:"column:poetry_id;comment:古诗id;size:11;"`
	Mp3      string `json:"mp3" form:"mp3" gorm:"column:mp3;comment:转换后的url;size:256;"`
}

func (PoetryLog) TableName() string {
	return "s_poetry_log"
}