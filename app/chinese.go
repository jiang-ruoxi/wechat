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

//ApiChineseBookList 国学绘本列表信息
func ApiChineseBookList(c *gin.Context) {
	////检查是否微信请求来源
	//if !common.CheckRequestUserAgent(c) {
	//	return
	//}
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	size := utils.GetIntParamItem("page_size", global.DEFAULT_PAGE_SIZE, c)
	level := utils.GetIntParamItem("level", global.DEFAULT_LEVEL, c)

	var service service.ChineseService
	list, total := service.GetChineseBookList(level, page, size)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"page_size":  size,
		"total_page": math.Ceil(float64(total) / float64(size)),
	}, global.SUCCESS_MSG, c)
}

//ApiChineseBookInfo 国学绘本详细信息
func ApiChineseBookInfo(c *gin.Context) {
	////检查是否微信请求来源
	//if !common.CheckRequestUserAgent(c) {
	//	return
	//}
	bookId := c.Query("book_id")
	var service service.ChineseService
	bookInfo := service.GetChineseBookInfo(bookId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, global.SUCCESS_MSG, c)
}
