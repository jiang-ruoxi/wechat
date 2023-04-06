package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wechat/common"
	"wechat/service"
	"wechat/utils"
)

func ApiIndex(c *gin.Context) {

	var service service.BaiKeService
	service.GetLPopData(1)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

//ApiQuestion 获取对应的栏目答题数据
func ApiQuestion(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Query("category_id"))
	var service service.BaiKeService
	question := service.GetLPopData(categoryId)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"data": question,
	}, common.SUCCESS_MSG, c)
}

//ApiAnswerList 获取答题记录
func ApiAnswerList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	var service service.BaiKeService
	list, total, err := service.GetAnswerList(page, pageSize)
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

//ApiAnswer 保存回答的答案
func ApiAnswer(c *gin.Context) {
	var req common.AnswerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"OpenId":      {utils.NotEmpty()},
		"CategoryId":  {utils.NotEmpty()},
		"QuestionId":  {utils.NotEmpty()},
		"IsSelect":    {utils.NotEmpty()},
		"RightSelect": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BaiKeService
	if err := service.InsertAnswer(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

//ApiLikeList 获取收藏记录
func ApiLikeList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	var service service.BaiKeService
	list, total, err := service.GetLikeList(page, pageSize)
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

//ApiLike 保存收藏的数据
func ApiLike(c *gin.Context) {
	var req common.LikeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"CategoryId": {utils.NotEmpty()},
		"QuestionId": {utils.NotEmpty()},
		"OpenId":     {utils.NotEmpty()},
		"Answer":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BaiKeService
	if err := service.InsertLike(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

//ApiUser 保存用户的数据
func ApiUser(c *gin.Context) {
	var req common.UserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"OpenId":   {utils.NotEmpty()},
		"NickName": {utils.NotEmpty()},
		"HeadUrl":  {utils.NotEmpty()},
		"Area":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BaiKeService
	if err := service.InsertUser(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}
