package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math"
	"net/http"
	"path"
	"strconv"
	"strings"
	"wechat/common"
	"wechat/global"
	"wechat/service"
	"wechat/utils"
)

func ApiPoetryList(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize < 1 {
		pageSize = 200
	}
	var service service.PoetryService
	list, total, err := service.GetPoetryList(page, pageSize)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list":       list,
		"page":       page,
		"total":      total,
		"page_size":  pageSize,
		"total_page": math.Ceil(float64(total) / float64(pageSize)),
	}, common.SUCCESS_MSG, c)
}

func ApiPoetryInfo(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var service service.PoetryService
	info := service.GetPoetryInfo(id)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"info": info,
	}, common.SUCCESS_MSG, c)
}

func GetPoetryOpenId(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	type OpenIdInfo struct {
		SessionKey string `json:"session_key"`
		Openid     string `json:"openid"`
	}
	var data OpenIdInfo
	code := c.Query("code")
	client := &http.Client{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=wx1090960cb624449c&secret=777c8da30ae1a4b152e386a896ce78a0&js_code=%s&grant_type=authorization_code", code)
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

func UploadPoetryMp3(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	file, err := c.FormFile("file")
	if err == nil {
		var Path string = "/data/web/static/poetry_log"
		dst := path.Join(Path, file.Filename)
		fmt.Printf("file.Filename:%s \n", file.Filename)
		fmt.Printf("dst:%s \n", dst)
		c.SaveUploadedFile(file, dst)
		dst = strings.Replace(dst, Path, "https://static.58haha.com/poetry_log", 1)
		fmt.Printf("dst:%s \n", dst)
		c.JSON(200, gin.H{
			"dst": dst,
		})
	} else {
		fmt.Println("UploadMp3失败")
		global.WECHAT_LOG.Info(fmt.Sprintf("AddUploads：%#v \n", err))
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
}

func AddPoetryVideoLog(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	var req common.PoetryVideoReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ReturnResponse(common.ERR_RES_PARAMS_ILLEGAL, map[string]interface{}{}, common.ERR_RES_PARAMS_ILLEGAL_MSG, c)
		return
	}
	verify := utils.Rules{
		"OpenId":   {utils.NotEmpty()},
		"PoetryId": {utils.NotEmpty()},
		"Mp3":      {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, err.Error(), c)
		return
	}
	var service service.PoetryService
	if err := service.InsertVideoLog(&req); err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
		return
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{}, common.SUCCESS_MSG, c)
}

func GetPoetryLog(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	poetryId, _ := strconv.Atoi(c.Query("poetry_id"))
	openId := c.Query("open_id")
	var service service.PoetryService
	info, total := service.GetPoetryLog(openId, poetryId)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"info":  info,
		"total": total,
	}, common.SUCCESS_MSG, c)
}

func ApiPoetryListCI(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize < 1 {
		pageSize = 1000
	}
	var service service.PoetryService
	list, total, err := service.GetPoetryListCI(page, pageSize)
	if err != nil {
		common.ReturnResponse(common.FAIL, map[string]interface{}{}, common.FAIL_MSG, c)
	}
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"list":       list,
		"page":       page,
		"total":      total,
		"page_size":  pageSize,
		"total_page": math.Ceil(float64(total) / float64(pageSize)),
	}, common.SUCCESS_MSG, c)
}

func ApiPoetryInfoCI(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var service service.PoetryService
	info := service.GetPoetryInfoCI(id)
	common.ReturnResponse(common.SUCCESS, map[string]interface{}{
		"info": info,
	}, common.SUCCESS_MSG, c)
}