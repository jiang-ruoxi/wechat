package model

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

// TableName Book 表名
func (ChengYU) TableName() string {
	return "s_chengyu"
}
