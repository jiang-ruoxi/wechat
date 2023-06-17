package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

func AddVideoLog(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	var req common.VideoLogReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"OpenId":   {utils.NotEmpty()},
		"BookId":   {utils.NotEmpty()},
		"Position": {utils.NotEmpty()},
		"Url":      {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BookService
	if err := service.InsertVideoLog(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

func UploadMp3(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	file, err := c.FormFile("file")
	if err == nil {
		var Path string = "/data/web/static/video"
		dst := path.Join(Path, file.Filename)
		fmt.Printf("file.Filename:%s \n", file.Filename)
		fmt.Printf("dst:%s \n", dst)
		c.SaveUploadedFile(file, dst)
		dst = strings.Replace(dst, Path, "https://static.58haha.com/video", 1)
		fmt.Printf("dst:%s \n", dst)
		c.JSON(200, gin.H{
			"dst": dst,
		})
	} else {
		fmt.Println("UploadMp3失败")
		global.WECHAT_LOG.Info(fmt.Sprintf("AddUploads：%#v \n", err))
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
}

func MakeVideo(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	var req common.VideoLogReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"OpenId":   {utils.NotEmpty()},
		"BookId":   {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BookService
	if err := service.MakeVideo(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}