package app

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

// ApiLexiconList 英语单词列表
func ApiLexiconList(c *gin.Context) {
	var service service.LexiconService
	page := utils.GetIntParamItem("point", global.DEFAULT_PAGE, c)
	lexiconType := utils.GetIntParamItem("type", global.DEFAULT_Type, c)
	size := utils.GetIntParamItem("size", global.DEFAULT_Lexicon_PAGE_SIZE, c)
	list, wordMap, AnnotationMap := service.ApiLexiconList(lexiconType, page, size)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"list":       list,
		"word":       wordMap,
		"annotation": AnnotationMap,
	}, global.SUCCESS_MSG, c)
}
