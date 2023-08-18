package router

import (
	"github.com/gin-gonic/gin"
	"wechat/app"
	"wechat/global"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	//路由组v1
	api := router.Group("v1")
	api.Use()
	{
		//首页demo
		api.GET("/index", app.ApiIndex)

		//古诗词成语
		//小学
		api.GET("/poetry/school/getList", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryList)
		api.GET("/poetry/school/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryInfo)
		//初高中
		api.GET("/poetry/junior/getList", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryList)
		api.GET("/poetry/junior/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryInfo)
		//成语
		api.GET("/poetry/cheng/getList", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryList)
		api.GET("/poetry/cheng/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryInfo)


		//中文绘本
		api.GET("/chinese/getList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookList)
		api.GET("/chinese/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookInfo)

		//英文绘本
		api.GET("/english/getList", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookList)
		api.GET("/english/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookInfo)

		//百科知识
	}
	return router
}
