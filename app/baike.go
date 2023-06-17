package app

import (
	"github.com/gin-gonic/gin"
	"strings"
	"wechat/common"
	"wechat/service"
	"wechat/utils"
)

// AddBaiKe 保存数据
func AddBaiKe(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
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

