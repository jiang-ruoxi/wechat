package service

import (
	"wechat/common/response"
	"wechat/global"
	"wechat/model"
	"regexp"
	"strings"
)

type PoetryService struct {
}

//GetSchoolPoetryList 小学古诗词列表
func (ps *PoetryService) GetSchoolPoetryList(page, size int) (poetryList []response.ResponseSchoolPoetry, total int64) {
	offset := size * (page - 1)
	db := global.GVA_DB.Model(&model.Poetry{}).Debug()
	db.Raw("SELECT id FROM s_poetry").Count(&total)
	db.Raw("SELECT id,poetry_id,title,grade_id,grade,grade_level,author,dynasty FROM s_poetry limit ? offset ?", size, offset).Scan(&poetryList)
	return
}

//GetSchoolPoetryInfo 小学古诗词详情
func (ps *PoetryService) GetSchoolPoetryInfo(poetryId int) (infoData response.ResponseSchoolPoetryData) {
	var info model.Poetry
	db := global.GVA_DB.Model(&model.Poetry{}).Debug()
	db = db.Where("poetry_id = ?", poetryId)
	db = db.Find(&info)

	str := strings.TrimPrefix(info.Content, "[")
	str = strings.TrimSuffix(str, "]")

	arr := strings.Split(str, "\",\"")
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.TrimSuffix(strings.TrimPrefix(arr[i], "\""), "\"")
	}

	str = info.Info
	re := regexp.MustCompile("[,.!\"]+")
	str = re.ReplaceAllString(str, " ")
	str = strings.TrimPrefix(str, "[[")
	str = strings.TrimSuffix(str, "]]")
	arr2d := [][]string{}
	for _, line := range strings.Split(str, "],") {
		line = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "["), "]"))
		arr2d = append(arr2d, strings.Split(line, ","))
	}
	for i, arr1d := range arr2d {
		for j, str := range arr1d {
			arr1d[j] = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\""))
		}
		arr2d[i] = arr1d
	}
	arr22 := strings.Split(arr2d[0][0], "] [ ")

	var PInfo response.PoetryInfo
	var poetryListInfo []response.PoetryInfo
	for _, item := range arr {
		PInfo.ZH = item
		poetryListInfo = append(poetryListInfo, PInfo)
	}
	for idx, _ := range poetryListInfo {
		poetryListInfo[idx].PY = arr22[idx]
	}

	infoData.ListContent = arr
	infoData.ListInfo = arr22
	infoData.Id = info.Id
	infoData.PoetryListInfo = poetryListInfo
	infoData.PoetryId = info.PoetryId
	infoData.Title = info.Title
	infoData.GradeId = info.GradeId
	infoData.Grade = info.Grade
	infoData.GradeLevel = info.GradeLevel
	infoData.Author = info.Author
	infoData.Dynasty = info.Dynasty
	infoData.Mp3 = info.Mp3
	infoData.Content = info.Content
	infoData.Info = info.Info

	return
}

//GetJuniorPoetryList 中学古诗词列表
func (ps *PoetryService) GetJuniorPoetryList(page, size int) (poetryList []response.ResponseSchoolPoetry, total int64) {
	offset := size * (page - 1)
	db := global.GVA_DB.Model(&model.JuniorPoetry{}).Debug()
	db.Raw("SELECT id FROM s_junior_poetry").Count(&total)
	db.Raw("SELECT id,poetry_id,title,grade_id,grade,author,dynasty FROM s_junior_poetry limit ? offset ?", size, offset).Scan(&poetryList)
	return
}

//GetJuniorPoetryInfo 中学古诗词详情
func (ps *PoetryService) GetJuniorPoetryInfo(poetryId int) (infoData response.ResponseSchoolPoetryData) {
	// 创建db
	var info model.JuniorPoetry
	db := global.GVA_DB.Model(&model.JuniorPoetry{}).Debug()
	db = db.Where("poetry_id = ?", poetryId)
	db = db.Find(&info)
	str := strings.TrimPrefix(info.Content, "[")
	str = strings.TrimSuffix(str, "]")

	arr := strings.Split(str, "\",\"")
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.TrimSuffix(strings.TrimPrefix(arr[i], "\""), "\"")
	}

	var PInfo response.PoetryInfo
	var poetryListInfo []response.PoetryInfo
	for _, item := range arr {
		PInfo.ZH = item
		poetryListInfo = append(poetryListInfo, PInfo)
	}

	infoData.PoetryListInfo = poetryListInfo
	infoData.Id = info.Id
	infoData.PoetryId = info.PoetryId
	infoData.Title = info.Title
	infoData.GradeId = info.GradeId
	infoData.Grade = info.Grade
	infoData.Author = info.Author
	infoData.Dynasty = info.Dynasty
	infoData.Content = info.Content

	return
}

//GetChengPoetryList 成语列表
func (ps *PoetryService) GetChengPoetryList(level, page, size int) (chengPoetryList []model.ChengYU, total int64) {
	offset := size * (page - 1)
	db := global.GVA_DB.Model(&model.ChengYU{}).Debug()
	db = db.Where("level = ?", level).Count(&total)
	db = db.Limit(size).Offset(offset).Find(&chengPoetryList)
	return
}

func (ps *PoetryService) ChengPoetryInfo(id int) (cy response.CYdATA) {
	// 创建db
	var cyInfo model.ChengYU
	db := global.GVA_DB.Model(&model.ChengYU{}).Debug()
	db = db.Where("id = ?", id).First(&cyInfo)
	fields := strings.Fields(cyInfo.Story)
	cy.Id = cyInfo.Id
	cy.Title = cyInfo.Title
	cy.Pinyin = cyInfo.Pinyin
	cy.Explain = cyInfo.Explain
	cy.Source = cyInfo.Source
	cy.Usage = cyInfo.Usage
	cy.Example = cyInfo.Example
	cy.Near = cyInfo.Near
	cy.Antonym = cyInfo.Antonym
	cy.Analyse = cyInfo.Analyse
	cy.Story = cyInfo.Story
	cy.Level = cyInfo.Level
	cy.StoryList = fields
	return
}
