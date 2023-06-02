package model

// VideoLog 结构体
type VideoLog struct {
	Id       int    `json:"id" form:"id" gorm:"column:id;comment:主键id;size:11;"`
	OpenId   string `json:"open_id" form:"open_id" gorm:"column:open_id;comment:用户open_id;size:128;"`
	BookId   int    `json:"book_id" form:"book_id" gorm:"column:book_id;comment:绘本id;size:11;"`
	Position uint   `json:"position" form:"position" gorm:"column:position;comment:排序位置;size:3;"`
	Url      string `json:"url" form:"url" gorm:"column:url;comment:转换后的url;size:256;"`
	AddTime  int64  `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间;size:20;"`
}

// TableName
func (VideoLog) TableName() string {
	return "s_huiben_video_log"
}
