package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"wechat/common"
	"wechat/model"
	"wechat/pkg/mysql"
	"wechat/pkg/redis"
	"wechat/utils"
)

type BaiKeService struct {
}

// 准备请求 body 数据结构
type body struct {
	Touser          string      `json:"touser"`
	TemplateID      string      `json:"template_id"`
	Page            string      `json:"page"`
	FormID          string      `json:"form_id"`
	Data            interface{} `json:"data"`
	EmphasisKeyword string      `json:"emphasis_keyword"`
}

// 准备 data 数据结构（可以根据实际需要修改）
type data struct {
	Keyword1 struct {
		Value string `json:"value"`
		Color string `json:"color"`
	} `json:"time5"`
	Keyword2 struct {
		Value string `json:"value"`
		Color string `json:"color"`
	} `json:"thing6"`
	Keyword3 struct {
		Value string `json:"value"`
		Color string `json:"color"`
	} `json:"thing8"`
	Keyword4 struct {
		Value string `json:"value"`
		Color string `json:"color"`
	} `json:"thing7"`
}

func (bs *BaiKeService) SendMsg(name string) {

	// 构建请求 body 数据
	dataJson := data{}
	dataJson.Keyword1.Value = utils.GetCurrentFormatDate()
	dataJson.Keyword1.Color = "#173177"
	dataJson.Keyword2.Value = "新增用户"
	dataJson.Keyword2.Color = "#173177"
	dataJson.Keyword3.Value = name + "，加入挑战百科知识小程序啦"
	dataJson.Keyword3.Color = "#173177"
	dataJson.Keyword4.Value = "百科知识公基"
	dataJson.Keyword4.Color = "#173177"
	requestBody := body{
		Touser:          "oqXuP4nEcrQdreKXPK7PpTQVXrbM",
		TemplateID:      "ts6dbTrgBPBjOAEB6FI6T_SzZqJfbQOxDvDGtozN9GU",
		Page:            "/pages/index/index",
		FormID:          "oqXuP4nEcrQdreKXPK7PpTQVXrbM" + strconv.FormatInt(time.Now().UnixMicro(), 10),
		Data:            dataJson,
		EmphasisKeyword: "keyword1.DATA",
	}

	// 将 body 数据转为 json
	requestBodyJson, err := json.Marshal(requestBody)
	log.Printf("请求参数:%#v \n", string(requestBodyJson))
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}

	//队列名称
	redisToken := "58haha_wechat_token"

	token, err := redis.RedisClient.Get(context.Background(), redisToken).Result()
	if err != nil {
		log.Printf("redis获取数据:%#v \n", err.Error())
	}
	log.Printf("redis获取token:%#v \n", token)
	if token == "" {
		token = bs.GetToken()
	}

	apiUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", token)
	log.Printf("请求参数url:%#v \n", apiUrl)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		fmt.Println("http.NewRequest error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do error:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应结果
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("json.NewDecoder error:", err)
		return
	}
	if result["errcode"].(float64) != 0 {
		fmt.Println("send template message error:", result["errmsg"].(string))
		return
	}
	fmt.Println("send template message success")
}

// PushDataToQueue 将对应栏目中的数据push到队列中
func (bs *BaiKeService) PushDataToQueue(categoryId int) error {
	// 创建db
	var baiKeList []model.BaiKe
	db := mysql.DB.Model(&model.BaiKe{}).Debug()

	if categoryId > 0 {
		db = db.Select("id").Where("category_id = ?", categoryId).Order("question desc, id desc").Find(&baiKeList)
	} else {
		db = db.Select("id").Order("question desc, id desc").Find(&baiKeList)
	}

	questionIds := make([]int, 0)
	for _, item := range baiKeList {
		questionIds = append(questionIds, item.Id)
	}

	//队列名称
	queue := fmt.Sprintf(common.DEFAULT_QUEUE)
	if categoryId > 0 {
		queue = fmt.Sprintf(common.QUEUE, categoryId)
	}

	pipe := redis.RedisClient.Pipeline()
	for _, item := range questionIds {
		pipe.RPush(context.Background(), queue, item)
	}
	if _, err := pipe.Exec(context.Background()); err != nil {
		return err
	}

	return nil
}

// GetToken 获取token
func (bs *BaiKeService) GetToken() string {

	//队列名称
	redisToken := "58haha_wechat_token"

	token, err := redis.RedisClient.Get(context.Background(), redisToken).Result()
	if err != nil {
		log.Printf("redis获取数据:%#v \n", err.Error())
	}

	if token != "" {
		return token
	}

	type TokenInfo struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   string `json:"expires_in"`
	}
	var tokenInfo TokenInfo

	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxb1e2d179618ef271&secret=02df528147a8ba1e5a4e3d9db537ee9d"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("http请求:%#v \n", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("body返回:%#v \n", err.Error())
	}
	json.Unmarshal(body, &tokenInfo)
	log.Printf("tokenInfo:%#v \n", tokenInfo)
	token = tokenInfo.AccessToken
	if token != "" {
		redis.RedisClient.Set(context.Background(), redisToken, token, 7000*time.Second)
		token, _ = redis.RedisClient.Get(context.Background(), redisToken).Result()
	}

	return token
}

// GetLPopData 从对应栏目中的队列中lpop数据
func (bs *BaiKeService) GetLPopData(categoryId int) map[string]interface{} {
	//获取之前判断队列中还有多少数据
	bs.IsCheckCount(categoryId)

	//队列名称
	queue := fmt.Sprintf(common.DEFAULT_QUEUE)
	if categoryId > 0 {
		queue = fmt.Sprintf(common.QUEUE, categoryId)
	}

	questionId, err := redis.RedisClient.LPop(context.Background(), queue).Result()
	if err != nil {
		fmt.Println("从队列中获取数据失败")
	}
	// 创建db
	var baiKe model.BaiKe
	db := mysql.DB.Model(&model.BaiKe{}).Debug()
	db = db.Where("id = ?", questionId).Find(&baiKe)

	return map[string]interface{}{
		"id":          baiKe.Id,
		"category_id": baiKe.CategoryId,
		"question":    baiKe.Question,
		"option_a":    baiKe.OptionA,
		"option_b":    baiKe.OptionB,
		"option_c":    baiKe.OptionC,
		"option_d":    baiKe.OptionD,
		"answer":      baiKe.Answer,
		"analytic":    baiKe.Analytic,
	}
}

// GetAnswerList 获取答题记录
func (bs *BaiKeService) GetAnswerList(page, pageSize int) (list []map[string]interface{}, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	openId := "oqXuP4nEcrQdreKXPK7PpTQVXrbM"
	// 创建db
	var answerList []model.Answer
	db := mysql.DB.Model(&model.Answer{}).Debug()
	db = db.Where("open_id = ?", openId)
	err = db.Count(&total).Error
	db = db.Order("add_time desc")
	db = db.Limit(limit).Offset(offset).Find(&answerList)

	//获取问题列表与栏目列表key=>value格式
	//questionIds := make([]string, 0)
	//for _, item := range answerList {
	//	//questionIds = append(questionIds, strconv.Itoa(item.QuestionId))
	//}
	//questionList := bs.GetQuestionKeyValue(questionIds)

	categoryIds := make([]string, 0)
	for _, item := range answerList {
		categoryIds = append(categoryIds, strconv.Itoa(item.Id))
	}
	//categoryList := bs.GetCategoryKeyValue(categoryIds)

	data := make([]map[string]interface{}, 0)

	for _, item := range answerList {
		answerTime, _ := strconv.ParseInt(item.AddTime, 10, 64)
		d := map[string]interface{}{
			"id":      item.Id,
			"open_id": item.OpenId,
			//"category":     categoryList[item.CategoryId],
			//"question":     questionList[item.QuestionId],
			"is_select":    item.IsSelect,
			"right_select": item.RightSelect,
			"answer_time":  utils.FormatDateFromUnix(answerTime),
		}
		data = append(data, d)
	}

	return data, total, err
}

// GetLikeList 获取答题记录
func (bs *BaiKeService) GetLikeList(page, pageSize int) (list []map[string]interface{}, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	openId := "oqXuP4nEcrQdreKXPK7PpTQVXrbM"
	// 创建db
	var likeList []model.Like
	db := mysql.DB.Model(&model.Like{}).Debug()
	db = db.Where("open_id = ?", openId)
	err = db.Count(&total).Error
	db = db.Order("add_time desc")
	db = db.Limit(limit).Offset(offset).Find(&likeList)

	//获取问题列表与栏目列表key=>value格式
	questionIds := make([]string, 0)
	for _, item := range likeList {
		questionIds = append(questionIds, strconv.Itoa(item.QuestionId))
	}
	questionList := bs.GetQuestionKeyValue(questionIds)
	answerList := bs.GetAnswerKeyValue(questionIds)

	categoryIds := make([]string, 0)
	for _, item := range likeList {
		categoryIds = append(categoryIds, strconv.Itoa(item.Id))
	}
	categoryList := bs.GetCategoryKeyValue(categoryIds)

	data := make([]map[string]interface{}, 0)

	for _, item := range likeList {
		answerTime, _ := strconv.ParseInt(item.AddTime, 10, 64)
		d := map[string]interface{}{
			"id":          item.Id,
			"open_id":     item.OpenId,
			"category":    categoryList[item.CategoryId],
			"question":    questionList[item.QuestionId],
			"answer":      answerList[item.QuestionId],
			"answer_time": utils.FormatDateFromUnix(answerTime),
		}
		data = append(data, d)
	}

	return data, total, err
}

// GetQuestionKeyValue 获取指定的问题
func (bs *BaiKeService) GetQuestionKeyValue(questionIds []string) map[int]string {
	var questionList []model.BaiKe
	db1 := mysql.DB.Model(&model.BaiKe{}).Debug()
	db1 = db1.Where("id in (?)", questionIds).Find(&questionList)
	question := make(map[int]string)
	for _, item := range questionList {
		question[item.Id] = item.Question
	}
	return question
}

// GetAnswerKeyValue 获取指定的问题的答案
func (bs *BaiKeService) GetAnswerKeyValue(questionIds []string) map[int]string {
	var questionList []model.BaiKe
	db1 := mysql.DB.Model(&model.BaiKe{}).Debug()
	db1 = db1.Where("id in (?)", questionIds).Find(&questionList)
	question := make(map[int]string)
	for _, item := range questionList {
		question[item.Id] = item.Answer
	}
	return question
}

// GetCategoryKeyValue 获取指定的栏目名称
func (bs *BaiKeService) GetCategoryKeyValue(categoryIds []string) map[int]string {
	var categoryList []model.Category
	db1 := mysql.DB.Model(&model.Category{}).Debug()
	db1 = db1.Where("id in (?)", categoryIds).Find(&categoryList)
	category := make(map[int]string)
	for _, item := range categoryList {
		category[item.Id] = item.Name
	}
	return category
}

// InsertLike 插入我的收藏
func (bs *BaiKeService) InsertLike(c *common.LikeReq) (err error) {
	//定义对应的类型
	var data model.Like
	//格式化数据生成
	c.GenerateLike(&data)
	if err = mysql.DB.Model(&model.Like{}).Create(&data).Error; err != nil {
		fmt.Println("数据创建失败")
		return err
	}
	return nil
}

// InsertUser 插入用户数据
func (bs *BaiKeService) InsertUser(c *common.UserReq) (err error) {
	//定义对应的类型
	var data model.User
	//格式化数据生成
	c.GenerateUser(&data)
	var count int64
	mysql.DB.Model(&model.User{}).Where("open_id = ?", data.OpenId).Count(&count)
	if count > 0 {
		return nil
	}
	if err = mysql.DB.Model(&model.User{}).Create(&data).Error; err != nil {
		fmt.Println("数据创建失败")
		return err
	}
	bs.SendMsg(data.NickName)
	return nil
}

// IsCheckCount 校验队列中的数据是否小于指定的数量
func (bs *BaiKeService) IsCheckCount(categoryId int) {
	//队列名称
	queue := common.DEFAULT_QUEUE
	if categoryId > 0 {
		queue = fmt.Sprintf(common.QUEUE, categoryId)
	}
	total, _ := redis.RedisClient.LLen(context.Background(), queue).Result()
	if total < common.QUEUE_LEN {
		bs.PushDataToQueue(categoryId)
	}
}

// DeleteQueue 删除指定的队列
func (bs *BaiKeService) DeleteQueue(categoryId int) {
	//队列名称
	queue := common.DEFAULT_QUEUE
	if categoryId > 0 {
		queue = fmt.Sprintf(common.QUEUE, categoryId)
	}
	redis.RedisClient.Del(context.Background(), queue).Result()
}

// GetInfoByOpenId 插入用户数据
func (bs *BaiKeService) GetInfoByOpenId(openId string) (data model.User, err error) {
	mysql.DB.Model(&model.User{}).Where("open_id = ?", openId).Find(&data)

	return data, nil
}

// GetCountByOpenId 插入用户数据
func (bs *BaiKeService) GetCountByOpenId(openId string) (count int64, err error) {
	mysql.DB.Model(&model.User{}).Where("open_id = ?", openId).Count(&count)

	return count, nil
}

// AddQuestion 插入用户数据
func (bs *BaiKeService) AddQuestion(openId, questionId, isSelect, rightSelect string) (err error) {
	var count int64
	mysql.DB.Model(&model.Answer{}).Where("open_id = ? and question_id = ?", openId, questionId).Count(&count)
	if count > 0 {
		return nil
	}
	var data model.Answer
	data.OpenId = openId
	data.QuestionId = questionId
	data.IsSelect = isSelect
	data.RightSelect = rightSelect
	data.CategoryId = "0"
	data.AddTime = "1111"
	if err = mysql.DB.Model(&model.Answer{}).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// InsertAnswer 插入答案数据
func (bs *BaiKeService) InsertAnswer(c *common.AnswerReq) (err error) {
	//定义对应的类型
	var data model.Answer
	//格式化数据生成
	c.GenerateAnswer(&data)
	if err = mysql.DB.Model(&model.Answer{}).Create(&data).Error; err != nil {
		fmt.Println("数据创建失败")
		return err
	}
	return nil
}

// SetScore 插入答案数据
func (bs *BaiKeService) SetScore(userId string, score string) (err error) {
	var data model.User
	mysql.DB.Model(&model.User{}).Where("open_id =?", userId).First(&data)
	scoreOld := data.Score
	scoreNew, _ := strconv.Atoi(score)
	scoreRes := scoreOld + scoreNew
	//格式化数据生成
	if err = mysql.DB.Model(&model.User{}).Where("open_id =?", userId).Update(map[string]interface{}{
		"score": scoreRes,
	}).Error; err != nil {
		fmt.Println("数据创建失败")
		return err
	}
	return nil
}

// GetRankList 插入答案数据
func (bs *BaiKeService) GetRankList() (rankMapList []map[string]interface{}, err error) {
	var data []model.User
	db := mysql.DB.Model(&model.User{}).Debug()
	err = db.Limit(100).Order("score desc,id desc").Find(&data).Error

	var rank int
	for _, item := range data {
		rankMap := make(map[string]interface{})
		rank = rank + 1
		rankMap["nick_name"] = item.NickName
		rankMap["head_url"] = item.HeadUrl
		rankMap["score"] = item.Score
		rankMap["rank"] = rank
		if item.Score < 1 {
			rankMap["score"] = "-"
			rankMap["rank"] = "-"
		}

		rankMapList = append(rankMapList, rankMap)

	}
	return rankMapList, err
}

// GetRank 插入答案数据
func (bs *BaiKeService) GetRank(userId string) (rankMap map[string]interface{}, err error) {
	var data []model.User
	db := mysql.DB.Model(&model.User{}).Debug()
	err = db.Limit(100).Order("score desc,id desc").Find(&data).Error

	var rank int
	dataMap := make(map[string]interface{})
	for _, item := range data {
		rank = rank + 1
		if item.OpenId == userId {
			dataMap["nick_name"] = item.NickName
			dataMap["head_url"] = item.HeadUrl
			dataMap["score"] = item.Score
			dataMap["rank"] = rank
			if item.Score < 1 {
				dataMap["score"] = "-"
				dataMap["rank"] = "-"
			}
		}
	}
	return dataMap, err
}
