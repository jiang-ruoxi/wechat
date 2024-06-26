package router

import (
	"github.com/gin-gonic/gin"
	"wechat/app"
	"wechat/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	//路由组v2 校验是否微信或者小程序请求访问
	apiV2 := router.Group("v2")
	apiV2.Use(middleware.CheckWechatMiddleware())
	{
		//获取栏目展示
		apiV2.GET("/getBookListNav", app.ApiBookNavList)
		//中文绘本
		apiV2.GET("/chinese/getList", app.ApiChineseBookList)
		apiV2.GET("/chinese/getBookInfo", app.ApiChineseBookInfo)
		//中文绘本专辑
		apiV2.GET("/chinese/getAlbumList", app.ApiChineseBookAlbumList)
		apiV2.GET("/chinese/getAlbumListInfo", app.ApiChineseBookAlbumListInfo)
		apiV2.GET("/chinese/getAlbumInfo", app.ApiChineseBookAlbumInfo)

		//古诗绘本
		apiV2.GET("/poetry/getList", app.ApiPoetryBookList1)
		apiV2.GET("/poetry/getBookInfo", app.ApiPoetryBookInfo)

		//英语绘本
		apiV2.GET("/english/getList", app.ApiEnglishBookList1)
		apiV2.GET("/english/getBookInfo", app.ApiEnglishBookInfo1)

		////中文绘本
		//apiV2.GET("/chinese/getList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookList)
		//apiV2.GET("/chinese/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookInfo)
		////中文绘本专辑
		//apiV2.GET("/chinese/getAlbumList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumList)
		//apiV2.GET("/chinese/getAlbumListInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumListInfo)
		//apiV2.GET("/chinese/getAlbumInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumInfo)
	}
	return router
}
