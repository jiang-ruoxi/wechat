package app

import (
	"github.com/gin-gonic/gin"
	"math"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//该文件为古诗词成语的api

//ApiSchoolPoetryList 小学古诗列表信息
func ApiSchoolPoetryList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	size := utils.GetIntParamItem("page_size", global.DEFAULT_PAGE_SIZE, c)

	var service service.PoetryService
	list, total := service.GetSchoolPoetryList(page, size)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"page_size":  size,
		"total_page": math.Ceil(float64(total) / float64(size)),
	}, global.SUCCESS_MSG, c)
}

//ApiSchoolPoetryInfo 小学古诗详细信息
func ApiSchoolPoetryInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_BOOK_ID, c)
	var service service.PoetryService
	bookInfo := service.GetSchoolPoetryInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

//ApiJuniorPoetryList 中学古诗列表信息
func ApiJuniorPoetryList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	size := utils.GetIntParamItem("page_size", global.DEFAULT_PAGE_SIZE, c)

	var service service.PoetryService
	list, total := service.GetJuniorPoetryList(page, size)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"page_size":  size,
		"total_page": math.Ceil(float64(total) / float64(size)),
	}, global.SUCCESS_MSG, c)
}

//ApiJuniorPoetryInfo 中学古诗详细信息
func ApiJuniorPoetryInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_BOOK_ID, c)
	var service service.PoetryService
	bookInfo := service.GetJuniorPoetryInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

//ApiChengPoetryList 成语列表信息
func ApiChengPoetryList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	size := utils.GetIntParamItem("page_size", global.DEFAULT_PAGE_SIZE, c)
	level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)

	var service service.PoetryService
	list, total := service.GetChengPoetryList(level, page, size)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"page_size":  size,
		"total_page": math.Ceil(float64(total) / float64(size)),
	}, global.SUCCESS_MSG, c)
}

//ApiChengPoetryInfo 成语详细信息
func ApiChengPoetryInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_BOOK_ID, c)
	var service service.PoetryService
	bookInfo := service.ChengPoetryInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}