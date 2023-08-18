package app

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"strconv"
)

func ApiIndex(c *gin.Context) {
	//检查是否微信请求来源
	if !common.CheckRequestUserAgent(c) {
		return
	}
	categoryId, _ := strconv.Atoi(c.Query("category_id"))
	var service service.IndexService
	service.ApiIndex(categoryId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{}, global.SUCCESS_MSG, c)
}
