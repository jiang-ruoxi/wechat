package app

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
	"wechat/service"
	"wechat/utils"
)

// AddBaiKe 保存数据
func AddBaiKe(c *gin.Context) {
	var req common.BaiKeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"CategoryId":   {utils.NotEmpty()},
		"Question": {utils.NotEmpty()},
		"OptionA":  {utils.NotEmpty()},
		"OptionB":     {utils.NotEmpty()},
		"Answer":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BaiKeService
	if err := service.InsertBaiKe(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

