package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"wechat/common"
	"wechat/service"
)

func GetMathItemList(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	count, _ := strconv.Atoi(c.Query("count"))
	max, _ := strconv.Atoi(c.Query("max"))
	//op 1加法，2减法，3混合
	op, _ := strconv.Atoi(c.Query("op"))

	var service service.MathService

	list := service.GetMathItemList(op, count, max, 1)

	common.ReturnResponse(common.SUCCESS, list, common.SUCCESS_MSG, c)
}
