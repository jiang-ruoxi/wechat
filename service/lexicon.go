package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"wechat/global"
	"wechat/model"
	"wechat/utils"
)

type LexiconService struct {
}

type LexiconResp struct {
	Id         int    `json:"id"`
	Word       string `json:"word"`
	Annotation string `json:"annotation"`
}

// ApiLexiconList 获取英语的列表信息
func (cs *LexiconService) ApiLexiconList(lexiconType int, page int, size int) (showQuestionList []string, word map[string]string, Annotation map[string]string) {
	lexiconSize := global.DEFAULT_Lexicon_PAGE_SIZE
	if size < lexiconSize {
		size = lexiconSize
	}
	offset := size * (page - 1)

	lexiconRespList := cs.getLexiconResp(size, offset, lexiconType)
	wordMap := make(map[string]string, 0)
	AnnotationMap := make(map[string]string, 0)
	for idx, _ := range lexiconRespList {
		wordMap[lexiconRespList[idx].Word] = lexiconRespList[idx].Annotation
		AnnotationMap[lexiconRespList[idx].Annotation] = lexiconRespList[idx].Word
		showQuestionList = append(showQuestionList, lexiconRespList[idx].Word)
		showQuestionList = append(showQuestionList, lexiconRespList[idx].Annotation)
	}

	return showQuestionList, wordMap, AnnotationMap
}

func (cs *LexiconService) ampersandString(ampersandsStr string) (ampersandList []string) {
	ampersands := utils.RemoveLettersAndAmpersands(ampersandsStr)
	split := strings.Split(ampersands, " ")
	for _, ampersand := range split {
		ampersand = utils.RemoveHashIfNeeded(ampersand)
		if ampersand != "" {
			ampersandList = append(ampersandList, ampersand)
		}
	}
	return
}

func (cs *LexiconService) getLexiconResp(size int, offset int, lexiconType int) (lexiconResp []LexiconResp) {

	var lexiconList []model.Lexicon
	LexiconDB := global.GVA_DB.Model(&model.Lexicon{}).Debug()
	LexiconDB = LexiconDB.Where("status=1 and type = ?", lexiconType)
	LexiconDB = LexiconDB.Order("id asc").Limit(size).Offset(offset)
	LexiconDB.Find(&lexiconList)

	var slice []string

	var lexiconRespItem LexiconResp
	for idx, _ := range lexiconList {
		explain := lexiconList[idx].Explain

		if err := json.Unmarshal([]byte(explain), &slice); err != nil {
			fmt.Println(err)
			return
		}
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(len(slice))

		lexiconRespItem.Id = lexiconList[idx].Id
		lexiconRespItem.Word = lexiconList[idx].Word
		lexiconRespItem.Annotation = slice[randomNumber]

		lexiconResp = append(lexiconResp, lexiconRespItem)
	}
	return
}

func (cs *LexiconService) shuffleStrings(slice []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(slice) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
