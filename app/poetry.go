package app

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"wechat/common"
	"wechat/service"
)

func ApiPoetryList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize < 1 {
		pageSize = 200
	}
	var service service.PoetryService
	list, total, err := service.GetPoetryList(page, pageSize)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list":       list,
		"page":       page,
		"total":      total,
		"page_size":  pageSize,
		"total_page": math.Ceil(float64(total) / float64(pageSize)),
	}, common.SUCCESS_MSG, c)
}

func ApiPoetryInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var service service.PoetryService
	info := service.GetPoetryInfo(id)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"info": info,
	}, common.SUCCESS_MSG, c)
}
