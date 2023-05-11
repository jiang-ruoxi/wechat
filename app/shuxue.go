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
	category := c.Query("category")

	var service service.ShuxueService
	list := service.GenerateAdditionList(category, count, max)

	common.ReturnResponse(common.SUCCESS, list, common.SUCCESS_MSG, c)
}
