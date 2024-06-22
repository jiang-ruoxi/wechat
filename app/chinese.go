package app

import (
	"github.com/gin-gonic/gin"
	"math"
	"wechat/common"
	"wechat/common/response"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//该文件为中文国学绘本的api

// ApiBookNavList 国学绘本Nav列表
func ApiBookNavList(c *gin.Context) {
	typeId := utils.GetIntParamItem("type", global.DEFAULT_PAGE, c)
	var service service.ChineseService
	list := service.ApiBookNavList(typeId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

// ApiChineseNavList 国学绘本Nav列表
func ApiChineseNavList(c *gin.Context) {
	var service service.ChineseService
	list := service.ApiChineseNavList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

// ApiChineseBookList 国学绘本列表信息
func ApiChineseBookList(c *gin.Context) {
	//page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	//level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)
	//
	//var service service.ChineseService
	//list, total := service.GetChineseBookList(level, page)
	var data response.ResponseChineseBook
	var list []response.ResponseChineseBook
	data.BookId = "784ff39684c7a23cfa09e77719256e20"
	data.Title = "侏罗纪的伙伴"
	data.Icon = "https://oss.58haha.com/chinese_book/cover/108.jpg"
	data.BookCount = "1"
	list = append(list, data)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      1,
		"page":       1,
		"total_page": math.Ceil(float64(1) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiChineseBookInfo 国学绘本详细信息
func ApiChineseBookInfo(c *gin.Context) {
	bookId := c.Query("book_id")
	var service service.ChineseService
	bookInfo := service.GetChineseBookInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

// ApiChineseBookAlbumList 国学绘本专辑列表信息
func ApiChineseBookAlbumList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)

	var service service.ChineseService
	list, total := service.GetChineseBookAlbumLList(page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiChineseBookAlbumListInfo 国学绘本专辑列表信息
func ApiChineseBookAlbumListInfo(c *gin.Context) {
	bookId := c.Query("book_id")
	var service service.ChineseService
	list, total := service.GetChineseBookAlbumListInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       1,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiChineseBookAlbumInfo 国学绘本专辑详细信息
func ApiChineseBookAlbumInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_PAGE, c)
	var service service.ChineseService
	bookInfo := service.GetChineseBookAlbumInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

// ApiPoetryBookList 古诗绘本列表信息
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

// ApiPoetryBookInfo 古诗绘本详细信息
func ApiPoetryBookInfo(c *gin.Context) {
	bookId := c.Query("book_id")
	var service service.ChineseService
	bookInfo := service.GetPoetryBookInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}
