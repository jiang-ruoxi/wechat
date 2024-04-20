package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"wechat/common"
	"wechat/common/request"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

//ApiDownLoadPic 下载图片
func ApiDownLoadPic(c *gin.Context) {
	var service service.DownLoadService
	page := utils.GetIntParamItem("page", global.DEFAULT_PAGE, c)
	service.GetDownLoadImages(page)
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{}, global.SUCCESS_MSG, c)
}

//ApiUploadFileData 上传图片
func ApiUploadFileData(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		common.ReturnResponse(global.FAIL, map[string]interface{}{}, global.FAIL_MSG, c)
		return
	}
	// 获取文件名
	fileName := filepath.Base(file.Filename)
	// 获取文件扩展名
	extension := filepath.Ext(fileName)

	// 生成文件名（使用时间戳）
	fName := strconv.FormatInt(time.Now().UnixNano(), 10) + extension

	path := "/data/static/pdf-img/"
	utils.ExistDir(path)

	// 保存文件到服务器
	c.SaveUploadedFile(file, filepath.Join(path, fName))

	dst := "https://oss.58haha.com/pdf-img/" + fName

	c.JSON(200, gin.H{
		"dst": dst,
	})
}

//ApiMakePdf 生成Pdf
func ApiMakePdf(c *gin.Context) {
	var json request.MakePDF
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var service service.PDF
	data, err := service.ApiMakePDF(json)
	msg := "success"
	if err != nil {
		msg = "fail"
	}
	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": data,
		"msg":  msg,
	}, global.SUCCESS_MSG, c)
}
