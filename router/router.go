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
		api.GET("/index/cache",routerCache(300), app.ApiIndex)
		api.GET("/index/cache1",routerCache(600), app.ApiIndex)

		//古诗词成语
		//小学
		api.GET("/poetry/school/getList", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryList)
		api.GET("/poetry/school/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiSchoolPoetryInfo)
		api.GET("/poetry/school/getOpenId", app.ApiSchoolOpenId)
		api.GET("/poetry/school/getPoetryLog", app.ApiPoetryLog)
		api.POST("/poetry/school/uploadMp3", app.ApiUploadPoetryMp3)
		api.POST("/poetry/school/addVideoLog", app.ApiAddPoetryVideoLog)
		//初高中
		api.GET("/poetry/junior/getList", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryList)
		api.GET("/poetry/junior/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiJuniorPoetryInfo)
		//成语
		api.GET("/poetry/cheng/getList", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryList)
		api.GET("/poetry/cheng/getPoetryInfo", routerCache(global.RedisURL_CACHE), app.ApiChengPoetryInfo)

		//中文绘本
		api.GET("/chinese/getList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookList)
		api.GET("/chinese/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookInfo)

		//古诗绘本
		api.GET("/poetry/book/getList", routerCache(global.RedisURL_CACHE), app.ApiPoetryBookList)
		api.GET("/poetry/book/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiPoetryBookInfo)

		//英文绘本
		api.GET("/english/getList", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookList)
		api.GET("/english/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookInfo)
		api.GET("/english/getOpenId", app.ApiOpenId)

		//百科知识
		api.GET("/baike/getCategoryCount", routerCache(global.RedisURL_CACHE), app.ApiCategoryCount)
		api.GET("/baike/getQuestion", app.ApiQuestion)

		//下载图片
		api.GET("/downLoad/pic", app.ApiDownLoadPic)
		api.GET("/http/post", app.ApiHttpPost)
	}
	return router
}
