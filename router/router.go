package router

import (
	"github.com/gin-gonic/gin"
	"wechat/app"
	"wechat/global"
	"wechat/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	//路由组v2 校验是否微信或者小程序请求访问
	apiV2 := router.Group("v2")
	apiV2.Use(middleware.CheckWechatMiddleware())
	{
		//获取栏目展示
		apiV2.GET("/getBookListNav", routerCache(global.RedisURL_CACHE), app.ApiBookNavList)
		//中文绘本
		apiV2.GET("/chinese/getList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookList)
		apiV2.GET("/chinese/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookInfo)
		//中文绘本专辑
		apiV2.GET("/chinese/getAlbumList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumList)
		apiV2.GET("/chinese/getAlbumListInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumListInfo)
		apiV2.GET("/chinese/getAlbumInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumInfo)

		//古诗绘本
		apiV2.GET("/poetry/getList", routerCache(global.RedisURL_CACHE), app.ApiPoetryBookList)
		apiV2.GET("/poetry/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiPoetryBookInfo)

		//英语绘本
		apiV2.GET("/english/getList", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookList)
		apiV2.GET("/english/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiEnglishBookInfo)

		//展示播放按钮
		apiV2.GET("/show/play", routerCache(global.RedisURL_CACHE), app.ApiShowPlay)

		////中文绘本
		//apiV2.GET("/chinese/getList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookList)
		//apiV2.GET("/chinese/getBookInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookInfo)
		////中文绘本专辑
		//apiV2.GET("/chinese/getAlbumList", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumList)
		//apiV2.GET("/chinese/getAlbumListInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumListInfo)
		//apiV2.GET("/chinese/getAlbumInfo", routerCache(global.RedisURL_CACHE), app.ApiChineseBookAlbumInfo)

		//古诗相关
		//获取朝代列表
		apiV2.GET("/poem/dynasty/list", routerCache(global.RedisURL_CACHE), app.ApiDynastyList)
		//引文列表
		apiV2.GET("/poem/quotes/list", routerCache(global.RedisURL_CACHE), app.ApiQuotesList)
		//集合类别
		apiV2.GET("/poem/kind/list", routerCache(global.RedisURL_CACHE), app.ApiKindList)
		//指定类别的集合
		apiV2.GET("/poem/collection/list", routerCache(global.RedisURL_CACHE), app.ApiCollectionList)
		//指定集合的作品列表
		apiV2.GET("/poem/collection/work/list", routerCache(global.RedisURL_CACHE), app.ApiCollectionWorkList)
		//获取古诗词详情
		apiV2.GET("/poem/info", routerCache(global.RedisURL_CACHE), app.ApiPoemInfo)
		//搜索古诗
		apiV2.GET("/poem/search", routerCache(global.RedisURL_CACHE), app.ApiPoemSearch)
		apiV2.GET("/poem/search/list", routerCache(global.RedisURL_CACHE), app.ApiPoemSearchList)
		//获取作者列表-朝代
		apiV2.GET("/poem/author/list", routerCache(global.RedisURL_CACHE), app.ApiAuthorList)
		//获取作者详情
		apiV2.GET("/poem/author/info", routerCache(global.RedisURL_CACHE), app.ApiAuthorInfo)
		//名言警句
		apiV2.GET("/poem/saying/list", routerCache(global.RedisURL_CACHE), app.ApiSayingList)
	}
	return router
}
