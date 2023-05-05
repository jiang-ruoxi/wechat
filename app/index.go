package app

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"

	"github.com/gin-gonic/gin"
)

func ApiIndex(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Query("category_id"))
	var service service.BaiKeService
	service.PushDataToQueue(categoryId)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

// ApiQuestion 获取对应的栏目答题数据
func ApiQuestion(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Query("category_id"))
	var service service.BaiKeService
	question := service.GetLPopData(categoryId)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"data": question,
	}, common.SUCCESS_MSG, c)
}

// ApiAnswerList 获取答题记录
func ApiAnswerList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	var service service.BaiKeService
	list, total, err := service.GetAnswerList(page, pageSize)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list":      list,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}, common.SUCCESS_MSG, c)
}

// ApiAnswer 保存回答的答案
func ApiAnswer(c *gin.Context) {
	var req common.AnswerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"OpenId":      {utils.NotEmpty()},
		"CategoryId":  {utils.NotEmpty()},
		"QuestionId":  {utils.NotEmpty()},
		"IsSelect":    {utils.NotEmpty()},
		"RightSelect": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BaiKeService
	if err := service.InsertAnswer(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

// ApiLikeList 获取收藏记录
func ApiLikeList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	var service service.BaiKeService
	list, total, err := service.GetLikeList(page, pageSize)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list":      list,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}, common.SUCCESS_MSG, c)
}

// ApiLike 保存收藏的数据
func ApiLike(c *gin.Context) {
	var req common.LikeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"CategoryId": {utils.NotEmpty()},
		"QuestionId": {utils.NotEmpty()},
		"OpenId":     {utils.NotEmpty()},
		"Answer":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BaiKeService
	if err := service.InsertLike(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

// ApiUser 保存用户的数据
func ApiUser(c *gin.Context) {
	var req common.UserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"OpenId":   {utils.NotEmpty()},
		"NickName": {utils.NotEmpty()},
		"HeadUrl":  {utils.NotEmpty()},
		"Area":     {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.BaiKeService
	if err := service.InsertUser(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

func ApiDeleteQueue(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Query("category_id"))
	var service service.BaiKeService
	service.DeleteQueue(categoryId)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

func GetOpenId(c *gin.Context) {
	type OpenIdInfo struct {
		SessionKey string `json:"session_key"`
		Openid     string `json:"openid"`
	}
	var data OpenIdInfo
	code := c.Query("code")
	client := &http.Client{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=wxb1e2d179618ef271&secret=02df528147a8ba1e5a4e3d9db537ee9d&js_code=%s&grant_type=authorization_code", code)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
	json.Unmarshal(body, &data)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"data": data.Openid,
	}, common.SUCCESS_MSG, c)
}

func AddUser(c *gin.Context) {
	var req common.UserReq
	openId := c.Query("openId")
	nickName := c.Query("nickName")
	avatarUrl := c.Query("avatarUrl")
	req.OpenId = openId
	req.NickName = nickName
	req.HeadUrl = avatarUrl
	var service service.BaiKeService
	if err := service.InsertUser(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

func GetInfoByOpenId(c *gin.Context) {
	openId := c.Query("openId")
	var service service.BaiKeService
	info, err := service.GetInfoByOpenId(openId)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"data": info,
	}, common.SUCCESS_MSG, c)
}

func GetCountByOpenId(c *gin.Context) {
	openId := c.Query("openId")
	var service service.BaiKeService
	count, err := service.GetCountByOpenId(openId)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"count": count,
	}, common.SUCCESS_MSG, c)
}

func AddQuestion(c *gin.Context) {
	openId := c.Query("openId")
	questionId := c.Query("questionId")
	isSelect := c.Query("isSelect")
	rightSelect := c.Query("rightSelect")

	var service service.BaiKeService
	err := service.AddQuestion(openId, questionId, isSelect, rightSelect)
	if err != nil {
		global.WECHAT_LOG.Info(fmt.Sprintf("AddQuestion：%#v \n", err))
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

func AddUploads(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {
		var Path string = "/data/web/static/wechat"
		dst := path.Join(Path, file.Filename)
		c.SaveUploadedFile(file, dst)
		dst = strings.Replace(dst, Path, "https://static.58haha.com/wechat", 1)
		c.JSON(200, gin.H{
			"dst": dst,
		})
	} else {
		global.WECHAT_LOG.Info(fmt.Sprintf("AddUploads：%#v \n", err))
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
}

func SetScore(c *gin.Context) {
	userId := c.Query("user_id")
	score := c.Query("score")
	var service service.BaiKeService
	err := service.SetScore(userId, score)
	if err != nil {
		global.WECHAT_LOG.Info(fmt.Sprintf("AddQuestion：%#v \n", err))
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

// GetRankList 获取答题记录
func GetRankList(c *gin.Context) {
	var service service.BaiKeService
	list, err := service.GetRankList()
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list": list,
	}, common.SUCCESS_MSG, c)
}

// GetRank 获取答题记录
func GetRank(c *gin.Context) {
	userId := c.Query("user_id")
	var service service.BaiKeService
	rank, err := service.GetRank(userId)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"rank": rank,
	}, common.SUCCESS_MSG, c)
}

// GetMsgVerify 获取答题记录
func GetMsgVerify(c *gin.Context) {
	const token = "58haha_wechat_token" // 微信小程序后台设置的 Token
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	// 将 token、timestamp、nonce 三个参数进行字典序排序
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	// 将三个参数字符串拼接成一个字符串进行 SHA1 加密
	str := strings.Join(strs, "")
	h := sha1.New()
	h.Write([]byte(str))
	hashedStr := hex.EncodeToString(h.Sum(nil))

	// 验证签名
	if hashedStr == signature {
		c.String(http.StatusOK, echostr)
	} else {
		c.String(http.StatusBadRequest, "Invalid signature")
	}
}
