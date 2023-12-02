package app

import (
	"github.com/gin-gonic/gin"
	"math"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//该文件为中文国学绘本的api

//ApiChineseNavList 国学绘本Nav列表
func ApiChineseNavList(c *gin.Context) {
	var service service.ChineseService
	list := service.ApiChineseNavList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

//ApiChineseBookList 国学绘本列表信息
func ApiChineseBookList(c *gin.Context) {
	//page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	//level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)
	//
	//var service service.ChineseService
	//list, total := service.GetChineseBookList(level, page)

	type dd struct {
		Id        int    `json:"-"`
		BookId    string `json:"book_id"`
		Title     string `json:"title"`
		Icon      string `json:"icon"`
		Level     uint8  `json:"-"`
		Position  uint8  `json:"-"`
		BookCount string `json:"book_count"`
	}
	var m dd
	m.BookId = "6c2665d7c3ed1e5bfd8ba600f026eb55"
	m.BookCount = "16"
	m.Icon = "https://static.58haha.com/chinese_book/cover/2021113086116408542418619.jpg"
	m.Title = "幸福的霸王龙·尊老爱幼"
	dds := make([]dd, 0)
	dds = append(dds, m)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": dds,
		//"total":      total,
		//"page":       page,
		//"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

//ApiChineseBookInfo 国学绘本详细信息
func ApiChineseBookInfo(c *gin.Context) {
	bookId := c.Query("book_id")
	var service service.ChineseService
	bookInfo := service.GetChineseBookInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

//ApiPoetryBookList 古诗绘本列表信息
func ApiPoetryBookList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)

	var service service.ChineseService
	list, total := service.GetPoetryBookList(level, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

//ApiPoetryBookInfo 古诗绘本详细信息
func ApiPoetryBookInfo(c *gin.Context) {
	bookId := c.Query("book_id")
	var service service.ChineseService
	bookInfo := service.GetPoetryBookInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}
