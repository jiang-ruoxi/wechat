package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"math"
	"strconv"
	"wechat/common"
	"wechat/service"
)

func ApiChineseBookList(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(uaText)
	fmt.Println("------获取---User-Agent---开始")
	fmt.Println(uaText)
	fmt.Printf("%v\n", ua.Mobile())  // => true
	name, version := ua.Browser()
	fmt.Printf("%v\n", name)    // => "Android"
	fmt.Printf("%v\n", version) // => "4.0"
	fmt.Println("------获取---User-Agent---结束")
	fmt.Println("------获取---HTTP_USER_AGENT---开始")
	uaText1 := c.Request.Header.Get("HTTP_USER_AGENT")
	fmt.Printf("%v\n", uaText1) // => "4.0"
	fmt.Println("------获取---HTTP_USER_AGENT---结束")
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
	var service service.ChineseBookService
	list, total, err := service.GetChineseBookList(level, page, pageSize)
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

func ApiChineseBookInfo(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Query("book_id"))
	var service service.ChineseBookService
	bookInfo := service.GetChineseBookInfo(bookId)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, common.SUCCESS_MSG, c)
}


func ApiChineseCYList(c *gin.Context) {
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
	var service service.ChineseBookService
	list, total, err := service.GetChineseCYList(level, page, pageSize)
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

func ApiChineseCYInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var service service.ChineseBookService
	bookInfo := service.GetChineseCYInfo(id)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"info": bookInfo,
	}, common.SUCCESS_MSG, c)
}