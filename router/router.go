package router

import (
	"time"
	"wechat/app"

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
	api.Use()
	{
		api.GET("/delete", app.ApiDeleteQueue)
		api.GET("/index", app.ApiIndex)
		api.GET("/question", app.ApiQuestion)
		api.GET("/answer_list", app.ApiAnswerList)
		api.POST("/answer", app.ApiAnswer)
		api.GET("/like_list", app.ApiLikeList)
		api.POST("/like", app.ApiLike)
		api.POST("/user", app.ApiUser)
		api.POST("/s_baike/addBaike", app.AddBaiKe)

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
		api.GET("/token", app.GetToken)
		api.GET("/send_msg", app.SendMsg)
		api.GET("/share", app.ShareInfo)
		api.GET("/getCategoryCount", app.GetCategoryCount)
		api.GET("/getListNumber", app.MakeNumerResult)
		api.GET("/getMathListNumber", app.GetMathItemList)

		api.GET("/math/chuzhong", app.ApiSXList)
		api.GET("/math/ssq", app.GetMathLottoList)
		api.GET("/huiben/token", app.ApiHBToken)
		api.GET("/en_book/getMiniList", app.ApiMiniList)
		api.GET("/en_book/getList", app.ApiBookList)
		api.GET("/en_book/sign/getList", app.ApiSignBookList)
		api.GET("/en_book/getBookInfo", app.ApiBookInfo)
		api.GET("/en_book/getEnBookOpenId", app.GetEnBookOpenId)
		api.POST("/en_book/addVideoLog", app.AddVideoLog)
		api.POST("/en_book/uploadMp3", app.UploadMp3)
		api.POST("/en_book/makeVideo", app.MakeVideo)

		api.GET("/s_poetry/getList", app.ApiPoetryList)
		api.GET("/s_poetry/getListCI", app.ApiPoetryListCI)
		api.GET("/s_poetry/getInfo", app.ApiPoetryInfo)
		api.GET("/s_poetry/getInfoCI", app.ApiPoetryInfoCI)
		api.GET("/s_poetry/getPoetryOpenId", app.GetPoetryOpenId)
		api.GET("/s_poetry/getPoetryLog", app.GetPoetryLog)
		api.POST("/s_poetry/uploadMp3", app.UploadPoetryMp3)
		api.POST("/s_poetry/addVideoLog", app.AddPoetryVideoLog)

		api.GET("/s_chinese_picture/getList", app.ApiChineseBookList)
		api.GET("/s_chinese_picture/getBookInfo", app.ApiChineseBookInfo)
		api.GET("/s_chengyu/getList", app.ApiChineseCYList)
		api.GET("/s_chengyu/getInfo", app.ApiChineseCYInfo)
	}

	return router
}
