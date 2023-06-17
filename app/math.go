package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wechat/common"
	"wechat/service"
)

func GetMathItemList(c *gin.Context) {

	count, _ := strconv.Atoi(c.Query("count"))
	max, _ := strconv.Atoi(c.Query("max"))
	//op 1加法，2减法，3混合
	op, _ := strconv.Atoi(c.Query("op"))

	var service service.MathService

	list := service.GetMathItemList(op, count, max, 1)

	common.ReturnResponse(common.SUCCESS, list, common.SUCCESS_MSG, c)
}
