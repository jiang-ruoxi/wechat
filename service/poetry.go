package service

import (
	"fmt"
	"regexp"
	"strings"
	"wechat/common"
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

type PoetryData struct {
	Id             int          `json:"id"`
	PoetryId       int          `json:"poetry_id"`
	Title          string       `json:"title"`
	GradeId        uint8        `json:"grade_id"`
	Grade          string       `json:"grade"`
	GradeLevel     uint8        `json:"grade_level"`
	Author         string       `json:"author"`
	Dynasty        string       `json:"dynasty" `
	Mp3            string       `json:"mp3" `
	Content        string       `json:"content"`
	Info           string       `json:"info"`
	ListContent    []string     `json:"list_content"`
	ListInfo       []string     `json:"list_info"`
	PoetryListInfo []PoetryInfo `json:"poetry_list_info"`
}

type PoetryInfo struct {
	ZH string `json:"zh"`
	PY string `json:"py"`
}

func (ps *PoetryService) GetPoetryInfo(poetryId int) (infoData PoetryData) {
	// 创建db
	var info model.Poetry
	db := mysql.DB.Model(&model.Poetry{}).Debug()
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
	fmt.Println(arr22)

	var PInfo PoetryInfo
	var poetryListInfo []PoetryInfo
	for _, item := range arr {
		PInfo.ZH = item
		poetryListInfo = append(poetryListInfo, PInfo)
	}
	for idx, _ := range poetryListInfo {
		poetryListInfo[idx].PY = arr22[idx]
	}

	var result PoetryData
	result.ListContent = arr
	result.ListInfo = arr22
	result.Id = info.Id
	result.PoetryListInfo = poetryListInfo
	result.PoetryId = info.PoetryId
	result.Title = info.Title
	result.GradeId = info.GradeId
	result.Grade = info.Grade
	result.GradeLevel = info.GradeLevel
	result.Author = info.Author
	result.Dynasty = info.Dynasty
	result.Mp3 = info.Mp3
	result.Content = info.Content
	result.Info = info.Info

	return result
}

func (ps *PoetryService) InsertVideoLog(c *common.PoetryVideoReq) (err error) {
	//定义对应的类型
	var data model.PoetryLog
	//格式化数据生成
	c.GeneratePoetryVideoLog(&data)

	var total int64
	db := mysql.DB.Model(&model.PoetryLog{}).Debug()
	db.Raw("SELECT count(id) as num FROM s_poetry_log where open_id = ? AND poetry_id = ?",data.OpenId, data.PoetryId).Count(&total)
	if total > 0 {
		mysql.DB.Model(&model.PoetryLog{}).Where("open_id = ? AND poetry_id = ?", data.OpenId, data.PoetryId).Delete(&model.PoetryLog{})
	}
	if err = mysql.DB.Model(&model.PoetryLog{}).Create(&data).Error; err != nil {
		fmt.Println("数据创建失败")
		return err
	}
	return nil
}