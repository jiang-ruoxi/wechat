package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"strconv"
	"wechat/global"
)

// GetIntParamItem 将获取的参数进行转成int类型
func GetIntParamItem(param string, defaultInt int, c *gin.Context) (paramInt int) {
	if paramInt, _ = strconv.Atoi(c.Query(param)); paramInt < global.DEFAULT_NUM {
		paramInt = defaultInt
	}
	return
}

//ExistDir 创建目录
func ExistDir(path string)  {
	// 判断路径是否存在
	_, err := os.ReadDir(path)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(path, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
}