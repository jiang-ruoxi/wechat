package service

import (
	"wechat/model"
	"wechat/pkg/mysql"
)

type PoetryService struct {
}

type PoetryInfoList struct {
	Id         int    `json:"id"`
	PoetryId   int    `json:"poetry_id"`
	Title      string `json:"title"`
	GradeId    uint8  `json:"grade_id" `
	Grade      string `json:"grade" `
	GradeLevel uint8  `json:"grade_level" `
	Author     string `json:"author" `
	Dynasty    string `json:"dynasty"`
}

// GetPoetryList 列表
func (ps *PoetryService) GetPoetryList(page, pageSize int) (poetryInfoList []PoetryInfoList, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)

	db := mysql.DB.Model(&model.Poetry{}).Debug()

	db.Raw("SELECT id FROM s_poetry").Count(&total)
	db.Raw("SELECT id,poetry_id,title,grade_id,grade,grade_level,author,dynasty FROM s_poetry limit ? offset ?", limit, offset).Scan(&poetryInfoList)

	return poetryInfoList, total, err
}

func (ps *PoetryService) GetPoetryInfo(poetryId int) (info model.Poetry) {
	// 创建db
	db := mysql.DB.Model(&model.Poetry{}).Debug()
	db = db.Where("poetry_id = ?", poetryId)
	db = db.Find(&info)

	return info
}
