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

// ApiCollectionWorkList 指定集合的作品列表
func ApiCollectionWorkList(c *gin.Context) {
	collectionId := utils.GetIntParamItem("collection_id", 1, c)
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	var service service.PoemService
	list, total := service.ApiCollectionWorkList(collectionId, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiPoemSearch 古诗词搜索
func ApiPoemSearch(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	tType := utils.GetIntParamItem("type", 1, c)
	value := c.Query("value")
	var service service.PoemService
	list, total := service.ApiPoemSearch(tType, page, value)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiPoemSearchList 古诗词搜索,多条件
func ApiPoemSearchList(c *gin.Context) {
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	title := c.Query("title")
	author := c.Query("author")
	dynasty := c.Query("dynasty")
	kindCn := c.Query("kind_cn")
	var service service.PoemService
	list, total := service.ApiPoemSearchList(title, author, dynasty, kindCn, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiPoemInfo 获取古诗词详情
func ApiPoemInfo(c *gin.Context) {
	workId := utils.GetIntParamItem("work_id", 0, c)
	var service service.PoemService
	info := service.ApiPoemInfo(workId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": info,
	}, global.SUCCESS_MSG, c)
}

// ApiAuthorInfo 获取作者详情
func ApiAuthorInfo(c *gin.Context) {
	authorId := utils.GetIntParamItem("author_id", 0, c)
	var service service.PoemService
	info := service.ApiAuthorInfo(authorId)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"info": info,
	}, global.SUCCESS_MSG, c)
}

// ApiAuthorList 获取作者列表
func ApiAuthorList(c *gin.Context) {
	dynasty := c.Query("dynasty")
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	var service service.PoemService
	list, total := service.ApiAuthorList(dynasty, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}

// ApiSayingList 名言警句
func ApiSayingList(c *gin.Context) {
	dynasty := c.Query("dynasty")
	author := c.Query("author")
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	var service service.PoemService
	list, total := service.ApiSayingList(dynasty, author, page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"total":      total,
		"page":       page,
		"total_page": math.Ceil(float64(total) / float64(global.DEFAULT_PAGE_SIZE)),
	}, global.SUCCESS_MSG, c)
}
