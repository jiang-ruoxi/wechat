package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"wechat/common"
	"wechat/global"
	"wechat/utils"
)

func ApiPoetryBookList(c *gin.Context) {
	page := utils.GetIntParamItem("page", 1, c)
	limit := utils.GetIntParamItem("limit", 50, c)
	typeId := utils.GetIntParamItem("type_id", 1, c)
	// 设置表单数据
	data := url.Values{
		"org_id":  {"0"},
		"user_id": {"1"},
		"diff":    {"all"},
		"is_read": {"all"},
		"sort":    {"id"},
		"limit":   {"" + strconv.Itoa(limit) + ""},
		"page":    {"" + strconv.Itoa(page) + ""},
		"type_id": {"" + strconv.Itoa(typeId) + ""},
	}

	// 将表单数据转换为字符串
	dataString := data.Encode()

	// 转换为字节流
	dataBytes := strings.NewReader(dataString)

	// 设置请求
	req, err := http.NewRequest("POST", "https://mzbook.com/api/book.Book/getBookList", dataBytes)
	if err != nil {
		panic(err)
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 打印响应
	var contentResult AutoGenerated
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": contentResult.Data,
	}, global.SUCCESS_MSG, c)
}

func ApiPoetryBookInfo(c *gin.Context) {
	id := utils.GetIntParamItem("id", 1, c)

	// 设置表单数据
	data := url.Values{
		"id": {"" + strconv.Itoa(id) + ""},
	}

	// 将表单数据转换为字符串
	dataString := data.Encode()

	// 转换为字节流
	dataBytes := strings.NewReader(dataString)

	// 设置请求
	req, err := http.NewRequest("POST", "https://mzbook.com/api/book.Book/getSentence", dataBytes)
	if err != nil {
		panic(err)
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 打印响应
	var contentResult AutoGeneratedResult
	if err := json.Unmarshal(body, &contentResult); err != nil {
		return
	}

	common.ReturnResponse(global.SUCCESS, map[string]interface{}{
		"data": contentResult.Data,
	}, global.SUCCESS_MSG, c)
}

type AutoGenerated struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		TypeID     string `json:"type_id"`
		Img        string `json:"img"`
		Difficulty string `json:"difficulty"`
		Author     string `json:"author"`
		FreeRank   string `json:"free_rank"`
		IsRead     int    `json:"is_read"`
		IsFree     int    `json:"is_free"`
		IsVip      int    `json:"is_vip"`
	} `json:"data"`
	TypeList struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Img  string `json:"img"`
	} `json:"type_list"`
}

type AutoGeneratedResult struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data map[string]Item `json:"data"`
}

type Item struct {
	ID           int         `json:"id"`
	BookID       int         `json:"book_id"`
	CnContent    string      `json:"cn_content"`
	EnContent    string      `json:"en_content"`
	SpellContent string      `json:"spell_content"`
	CnArr        string      `json:"cn_arr"`
	EnArr        string      `json:"en_arr"`
	PlayTime     int         `json:"play_time"`
	EnPlayTime   int         `json:"en_play_time"`
	Img          string      `json:"img"`
	Meaning      string      `json:"meaning"`
	MeaningCnArr interface{} `json:"meaning_cn_arr"`
	MeaningEn    string      `json:"meaning_en"`
	MeaningEnArr interface{} `json:"meaning_en_arr"`
	Rank         int         `json:"rank"`
	IsShow       int         `json:"is_show"`
}
