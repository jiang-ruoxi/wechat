package middleware

import (
	"github.com/gin-gonic/gin"
	"wechat/common"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 鉴权取头部信息 x-token 信息
		token := c.Request.Header.Get("x-token")
		if len(token) < 1 || token != common.AUTH_TOKEN {
			common.ReturnResponse(403, map[string]interface{}{}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
