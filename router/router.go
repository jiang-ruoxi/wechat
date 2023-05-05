package router

import (
	"time"
	"wechat/app"
	"wechat/middleware"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	// 使用zap日志库
	router.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(zap.L(), true))
	//后台路由组
	api := router.Group("v1")
	api.Use(middleware.Auth())
	{
		api.GET("/delete", app.ApiDeleteQueue)
		api.GET("/index", app.ApiIndex)
		api.GET("/question", app.ApiQuestion)
		api.GET("/answer_list", app.ApiAnswerList)
		api.POST("/answer", app.ApiAnswer)
		api.GET("/like_list", app.ApiLikeList)
		api.POST("/like", app.ApiLike)
		api.POST("/user", app.ApiUser)

		api.GET("/getOpenId", app.GetOpenId)
		api.GET("/addUser", app.AddUser)
		api.GET("/getCountByOpenId", app.GetCountByOpenId)
		api.GET("/getInfoByOpenId", app.GetInfoByOpenId)
		api.GET("/addQuestion", app.AddQuestion)
		api.GET("/setScore", app.SetScore)
		api.GET("/getRankList", app.GetRankList)
		api.GET("/getRank", app.GetRank)
		api.POST("/upload", app.AddUploads)

		api.GET("/msg/verify", app.GetMsgVerify)

	}

	return router
}
