package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"wechat/common/request"
	"wechat/common/response"
	"wechat/global"
	"wechat/model"
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

//GetSchoolOpenId 获取open_id信息
func (ps *PoetryService) GetSchoolOpenId(code string) (openId string) {
	var data response.OpenIdData
	appid := global.GVA_CONFIG.Wechat.SchoolAppId
	secret := global.GVA_CONFIG.Wechat.SchoolSecret
	client := &http.Client{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, code)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	openId = data.Openid
	return
}

//GetPoetryLog
func (ps *PoetryService) GetPoetryLog(openId string, poetryId int) (infoData model.PoetryLog, total int64) {
	// 创建db
	var info model.PoetryLog
	db := global.GVA_DB.Model(&model.PoetryLog{}).Debug()
	db = db.Where("open_id = ? and poetry_id = ?", openId, poetryId).Count(&total)
	db = db.Find(&info)
	return info, total
}

//InsertVideoLog
func (ps *PoetryService) InsertVideoLog(c *request.PoetryVideoReq) (err error) {
	//定义对应的类型
	var data model.PoetryLog
	//格式化数据生成
	c.GeneratePoetryVideoLog(&data)

	var total int64
	db := global.GVA_DB.Model(&model.PoetryLog{}).Debug()
	db.Raw("SELECT count(id) as num FROM s_poetry_log where open_id = ? AND poetry_id = ?", data.OpenId, data.PoetryId).Count(&total)
	if total > 0 {
		global.GVA_DB.Model(&model.PoetryLog{}).Where("open_id = ? AND poetry_id = ?", data.OpenId, data.PoetryId).Delete(&model.PoetryLog{})
	}
	if err = global.GVA_DB.Model(&model.PoetryLog{}).Create(&data).Error; err != nil {
		return err
	}
	return nil
}
