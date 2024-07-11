package app

import (
	"github.com/gin-gonic/gin"
	"math"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

// ApiDynastyList 古诗词朝代列表
func ApiDynastyList(c *gin.Context) {
	var service service.PoemService
	list := service.ApiDynastyList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

// ApiQuotesList 古诗词引文列表
func ApiQuotesList(c *gin.Context) {
	dynasty := c.Query("dynasty")
	kind := c.Query("kind")
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	var service service.PoemService
	list, total := service.ApiQuotesList(dynasty, kind, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiKindList 古诗词集合类别列表
func ApiKindList(c *gin.Context) {
	var service service.PoemService
	list := service.ApiKindList()
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}

// ApiCollectionList 指定类别的集合
func ApiCollectionList(c *gin.Context) {
	kindId := utils.GetIntParamItem("kind_id", global.DEFAULT_PAGE, c)
	var service service.PoemService
	list := service.ApiCollectionList(kindId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list": list,
	}, global.SUCCESS_MSG, c)
}
