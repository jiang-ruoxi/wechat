package common

import (
	"strconv"
	"time"
	"wechat/model"
)

type LikeReq struct {
	Id         int    `json:"id" comment:""`
	OpenId     string `json:"open_id" comment:"open_id"`
	CategoryId int    `json:"category_id" comment:"category_id"`
	QuestionId int    `json:"question_id" comment:"question_id"`
	Answer     string `json:"answer" comment:"answer"`
	AddTime    string `json:"add_time" comment:"add_time"`
}

type AnswerReq struct {
	Id          int    `json:"id" comment:""`
	OpenId      string `json:"open_id" comment:"open_id"`
	CategoryId  int    `json:"category_id" comment:"category_id"`
	QuestionId  int    `json:"question_id" comment:"question_id"`
	IsSelect    string `json:"is_select" comment:"is_select"`
	RightSelect string `json:"right_select" comment:"right_select"`
	AddTime     string `json:"add_time" comment:"add_time"`
}

type UserReq struct {
	Id       int    `json:"id" comment:""`
	OpenId   string `json:"open_id" comment:"open_id"`
	NickName string `json:"nick_name" comment:"nick_name"`
	HeadUrl  string `json:"head_url" comment:"head_url"`
	Area     string `json:"area" comment:"area"`
	AddTime  string `json:"add_time" comment:"add_time"`
}

type VideoLogReq struct {
	OpenId   string `json:"open_id" comment:"open_id"`
	BookId   int    `json:"book_id" comment:"book_id"`
	Position uint   `json:"position" comment:"position"`
	Url      string `json:"url" comment:"url"`
}

func (req LikeReq) GenerateLike(model *model.Like) {
	model.OpenId = req.OpenId
	model.CategoryId = req.CategoryId
	model.QuestionId = req.QuestionId
	model.Answer = req.Answer
	model.AddTime = strconv.Itoa(int(time.Now().UTC().Unix()))
}

func (req AnswerReq) GenerateAnswer(model *model.Answer) {
	model.OpenId = req.OpenId
	model.IsSelect = req.IsSelect
	model.RightSelect = req.RightSelect
	model.AddTime = strconv.Itoa(int(time.Now().UTC().Unix()))
}

func (req UserReq) GenerateUser(model *model.User) {
	model.OpenId = req.OpenId
	model.NickName = req.NickName
	model.HeadUrl = req.HeadUrl
	model.Area = req.Area
	model.AddTime = strconv.Itoa(int(time.Now().UTC().Unix()))
}

func (req VideoLogReq) GenerateVideoLog(model *model.VideoLog) {
	model.OpenId = req.OpenId
	model.BookId = req.BookId
	model.Position = req.Position
	model.Url = req.Url
	model.AddTime = int64(time.Now().UTC().Unix())
}
