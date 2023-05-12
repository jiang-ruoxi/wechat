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
	op := c.Query("op")

	var service service.ShuxueService

	list := service.GenerateAdditionList(op, count, max)

	common.ReturnResponse(common.SUCCESS, list, common.SUCCESS_MSG, c)
}
