package app

import (
	"math"
	"strconv"
	"wechat/common"
	"wechat/service"

	"github.com/gin-gonic/gin"
)

func MakeNumerResult(c *gin.Context) {

	count, _ := strconv.Atoi(c.Query("count"))
	max, _ := strconv.Atoi(c.Query("max"))
	//op 1加法，2减法，3混合
	op := c.Query("op")
	//et 1加法不进位，2减法不退位 3加法进位 4减法退位
	et := c.Query("et")

	var service service.ShuxueService

	list := service.GenerateAdditionList(op, count, max, et)

	common.ReturnResponse(common.SUCCESS, list, common.SUCCESS_MSG, c)
}

// ApiSXList 数学列表
func ApiSXList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	var service service.ShuxueService
	list, total, err := service.GetSXList(page, pageSize)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list":      list,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}, common.SUCCESS_MSG, c)
}

// ApiHBToken 绘本token
func ApiHBToken(c *gin.Context) {
	var service service.ShuxueService
	token := service.GetHBToken()
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"token": token,
	}, common.SUCCESS_MSG, c)
}

func ApiBookList(c *gin.Context) {
	sign := c.Query("t")

	var service service.BookService
	isSign := service.IsCheckSign(sign)
	if !isSign {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize < 1 {
		pageSize = 18
	}
	level, _ := strconv.Atoi(c.Query("level"))
	if level < 1 {
		level = 1
	}

	list, total, err := service.GetBookList(level, page, pageSize)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_page": math.Ceil(float64(total) / float64(pageSize)),
	}, common.SUCCESS_MSG, c)
}

func ApiBookInfo(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Query("book_id"))
	var service service.BookService
	bookInfo := service.GetBookInfo(bookId)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, common.SUCCESS_MSG, c)
}
