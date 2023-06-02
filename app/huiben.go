package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
	"wechat/common"
	"wechat/global"
)

func AddVideoLog(c *gin.Context) {

}

func UploadMp3(c *gin.Context)  {
	file, err := c.FormFile("file")
	global.WECHAT_LOG.Info(fmt.Sprintf("上传录音接收到的参数：%#v \n", file))
	if err == nil {
		var Path string = "/data/web/static/video"
		dst := path.Join(Path, file.Filename)
		c.SaveUploadedFile(file, dst)
		dst = strings.Replace(dst, Path, "https://static.58haha.com/video", 1)
		c.JSON(200, gin.H{
			"dst": dst,
		})
	} else {
		global.WECHAT_LOG.Info(fmt.Sprintf("AddUploads：%#v \n", err))
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
}