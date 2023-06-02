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
	if err == nil {
		var Path string = "/data/web/static/video"
		dst := path.Join(Path, file.Filename)
		fmt.Printf("file.Filename:%s \n",file.Filename)
		fmt.Printf("dst:%s \n",dst)
		c.SaveUploadedFile(file, dst)
		dst = strings.Replace(dst, Path, "https://static.58haha.com/video", 1)
		fmt.Printf("dst:%s \n",dst)
		c.JSON(200, gin.H{
			"dst": dst,
		})
	} else {
		fmt.Println("UploadMp3失败")
		global.WECHAT_LOG.Info(fmt.Sprintf("AddUploads：%#v \n", err))
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
}