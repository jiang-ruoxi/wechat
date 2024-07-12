package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"
	"wechat/global"
)

// GetIntParamItem 将获取的参数进行转成int类型
func GetIntParamItem(param string, defaultInt int, c *gin.Context) (paramInt int) {
	if paramInt, _ = strconv.Atoi(c.Query(param)); paramInt < global.DEFAULT_NUM {
		paramInt = defaultInt
	}
	return
}

// ExistDir 创建目录
func ExistDir(path string) {
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

func ReplaceURLPart(originalURL, oldPart, newPart string) string {
	// 使用正则表达式匹配URL中的某个部分
	reg := regexp.MustCompile(regexp.QuoteMeta(oldPart))

	// 替换匹配到的部分
	return reg.ReplaceAllString(originalURL, newPart)
}

func RemoveLettersAndAmpersands(s string) string {
	reg := regexp.MustCompile("[a-zA-Z&]")
	return reg.ReplaceAllString(s, "")
}

func RemoveHashIfNeeded(s string) string {
	if strings.HasPrefix(s, ".") {
		return strings.TrimPrefix(s, ".")
	}
	return s
}

func SliceUnique(s []any) []any {
	m := make(map[any]bool)
	for _, v := range s {
		if !m[v] {
			m[v] = true
		}
	}
	result := make([]any, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}
