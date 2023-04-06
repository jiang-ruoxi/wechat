package router

import (
	"github.com/gin-gonic/gin"
	"wechat/app"
	"wechat/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

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
	}

	return router
}
