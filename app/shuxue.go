package app

import (
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