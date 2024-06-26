package app

import (
	"github.com/gin-gonic/gin"
	"math"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

func ApiBookNavList(c *gin.Context) {
	typeId := utils.GetIntParamItem("type", global.DEFAULT_PAGE, c)
	typeId = 101
	var service service.ChineseService
	list := service.ApiBookNavList(typeId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

func ApiChineseBookList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)

	level = 101
	var service service.ChineseService
	list, total := service.GetChineseBookList(level, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

func ApiChineseBookInfo(c *gin.Context) {
	bookId := c.Query("book_id")
	var service service.ChineseService
	bookInfo := service.GetChineseBookInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}

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

func ApiChineseBookAlbumInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", global.DEFAULT_PAGE, c)
	var service service.ChineseService
	bookInfo := service.GetChineseBookAlbumInfo(id)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}
