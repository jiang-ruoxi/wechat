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
	BookId   string `json:"book_id" comment:"book_id"`
	Position uint   `json:"position" comment:"position"`
	Url      string `json:"url" comment:"url"`
}

type PoetryVideoReq struct {
	OpenId   string `json:"open_id" comment:"open_id"`
	PoetryId int    `json:"poetry_id" comment:"poetry_id"`
	Mp3      string `json:"mp3" comment:"mp3"`
}

type BaiKeReq struct {
	Id         int    `json:"id" comment:"id"`
	CategoryId int    `json:"category_id" comment:"category_id"`
	Question   string `json:"question" comment:"question"`
	OptionA    string `json:"option_a" comment:"option_a"`
	OptionB    string `json:"option_b" comment:"option_b"`
	OptionC    string `json:"option_c" comment:"option_c"`
	OptionD    string `json:"option_d" comment:"option_d"`
	Answer     string `json:"answer" comment:"answer"`
	Analytic   string `json:"analytic" comment:"analytic"`
	AddTime    string `json:"add_time" comment:"add_time"`
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

func (req PoetryVideoReq) GeneratePoetryVideoLog(model *model.PoetryLog) {
	model.OpenId = req.OpenId
	model.PoetryId = req.PoetryId
	model.Mp3 = req.Mp3
}

func (req BaiKeReq) GenerateBaiKe(model *model.BaiKe) {
	model.CategoryId = req.CategoryId
	model.Question = req.Question
	model.OptionA = req.OptionA
	model.OptionB = req.OptionB
	model.OptionC = req.OptionC
	model.OptionD = req.OptionD
	model.Answer = req.Answer
	model.Analytic = req.Analytic
	model.AddTime = strconv.Itoa(int(time.Now().UTC().Unix()))
}
