package service

import (
	"strings"
	"wechat/global"
	"wechat/model"
	"wechat/utils"
)

type PoemService struct {
}

type DynastyResponse struct {
	DynastyId   int    `json:"dynasty_id"`
	DynastyName string `json:"dynasty_name"`
}

// ApiDynastyList 古诗词朝代列表
func (ps *PoemService) ApiDynastyList() (dynastyList []DynastyResponse) {
	var dynastyModelList []model.RXDynasty
	bookDB := global.GVA_DB.Model(&model.RXDynasty{}).Debug()
	bookDB = bookDB.Order("s_sort desc").Order("id asc")
	bookDB.Find(&dynastyModelList)
	var dynastyTemp DynastyResponse
	for idx, _ := range dynastyModelList {
		dynastyTemp.DynastyId = dynastyModelList[idx].DynastyId
		dynastyTemp.DynastyName = dynastyModelList[idx].Name
		dynastyList = append(dynastyList, dynastyTemp)
	}
	return
}

type QuoteResponse struct {
	QuoteId   int    `json:"quote_id"`
	Quote     string `json:"quote"`
	Dynasty   string `json:"dynasty"`
	Author    string `json:"author"`
	Kind      string `json:"kind"`
	WorkId    int    `json:"work_id"`
	WorkTitle string `json:"work_title"`
}

// ApiQuotesList 古诗词引文列表
func (ps *PoemService) ApiQuotesList(dynasty, kind string, page int) (quoteList []QuoteResponse, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	var quotesModelList []model.RXQuotes
	db := global.GVA_DB.Model(&model.RXQuotes{}).Debug()
	if dynasty != "" {
		db = db.Where("dynasty = ?", dynasty)
	}
	if kind != "" {
		db = db.Where("kind = ?", kind)
	}
	db = db.Count(&total)
	db = db.Limit(size).Offset(offset)
	db.Find(&quotesModelList)

	var quoteTemp QuoteResponse
	for idx, _ := range quotesModelList {
		quoteTemp.QuoteId = quotesModelList[idx].QuoteId
		quoteTemp.Quote = quotesModelList[idx].Quote
		quoteTemp.Dynasty = quotesModelList[idx].Dynasty
		quoteTemp.Author = quotesModelList[idx].Author
		quoteTemp.Kind = quotesModelList[idx].Kind
		quoteTemp.WorkId = quotesModelList[idx].WorkId
		quoteTemp.WorkTitle = quotesModelList[idx].WorkTitle
		quoteList = append(quoteList, quoteTemp)
	}
	return
}

type KindResponse struct {
	Id     int    `json:"id"`
	KindId int    `json:"kind_id"`
	Kind   string `json:"kind"`
	Name   string `json:"name"`
}

// ApiKindList 古诗词集合类别列表
func (ps *PoemService) ApiKindList(page, kindId int) (KindList []KindResponse, total int64) {
	size := global.DEFAULT_PAGE_SIZE_MAX
	offset := size * (page - 1)
	var collectionModelList []model.RXCollections
	db := global.GVA_DB.Model(&model.RXCollections{}).Debug()
	if kindId > 0 {
		db = db.Where("kind_id = ?", kindId)
	}
	db = db.Select("id,`name`,kind,kind_id")
	db = db.Order("id ASC")
	db.Group("kind_id,`name`,id")
	db = db.Count(&total)
	db = db.Limit(size).Offset(offset)
	db.Find(&collectionModelList)
	var kindTemp KindResponse
	for idx, _ := range collectionModelList {
		kindTemp.Id = collectionModelList[idx].Id
		kindTemp.KindId = collectionModelList[idx].KindId
		kindTemp.Kind = collectionModelList[idx].Kind
		kindTemp.Name = collectionModelList[idx].Name
		KindList = append(KindList, kindTemp)
	}
	return
}

type CollectionResponse struct {
	CollectionId   int    `json:"collection_id"`
	CollectionName string `json:"collection_name"`
	Cover          string `json:"cover"`
	QuotesCount    int    `json:"quotes_count"`
	WorksCount     int    `json:"works_count"`
}

// ApiCollectionList 指定类别的集合
func (ps *PoemService) ApiCollectionList(kindId int) (collectionList []CollectionResponse) {
	var collectionModelList []model.RXCollections
	db := global.GVA_DB.Model(&model.RXCollections{}).Debug().Where("kind_id = ?", kindId)
	db = db.Order("sort asc")
	db.Find(&collectionModelList)

	var collectionTemp CollectionResponse
	for idx, _ := range collectionModelList {
		collectionTemp.CollectionId = collectionModelList[idx].CollectionId
		collectionTemp.CollectionName = collectionModelList[idx].Name
		collectionTemp.Cover = collectionModelList[idx].Cover
		collectionTemp.QuotesCount = collectionModelList[idx].QuotesCount
		collectionTemp.WorksCount = collectionModelList[idx].WorksCount
		collectionList = append(collectionList, collectionTemp)
	}
	return
}

type CollectionWorkResponse struct {
	CollectionWorkId int    `json:"collection_work_id"`
	WorkId           int    `json:"work_id"`
	CollectionId     int    `json:"collection_id"`
	WorkTitle        string `json:"work_title"`
	WorkAuthor       string `json:"work_author"`
	WorkDynasty      string `json:"work_dynasty"`
	WorkContent      string `json:"work_content"`
	WorkKind         string `json:"work_kind"`
	Collection       string `json:"collection"`
}

// ApiCollectionWorkList 指定集合的作品列表
func (ps *PoemService) ApiCollectionWorkList(collectionId, page int) (collectionWorkList []CollectionWorkResponse, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	var collectionWorkModelList []model.RXCollectionWorks
	db := global.GVA_DB.Model(&model.RXCollectionWorks{}).Debug().Where("collection_id = ?", collectionId)
	db = db.Order("sort asc")
	db = db.Count(&total)
	db = db.Limit(size).Offset(offset)
	db.Find(&collectionWorkModelList)

	var collectionWorkTemp CollectionWorkResponse
	for idx, _ := range collectionWorkModelList {
		collectionWorkTemp.CollectionWorkId = collectionWorkModelList[idx].CollectionId
		collectionWorkTemp.WorkId = collectionWorkModelList[idx].WorkId
		collectionWorkTemp.CollectionId = collectionWorkModelList[idx].CollectionId
		collectionWorkTemp.WorkTitle = collectionWorkModelList[idx].WorkTitle
		collectionWorkTemp.WorkAuthor = collectionWorkModelList[idx].WorkAuthor
		collectionWorkTemp.WorkDynasty = collectionWorkModelList[idx].WorkDynasty
		collectionWorkTemp.WorkContent = collectionWorkModelList[idx].WorkContent
		collectionWorkTemp.WorkKind = collectionWorkModelList[idx].WorkKind
		collectionWorkTemp.Collection = collectionWorkModelList[idx].Collection
		collectionWorkList = append(collectionWorkList, collectionWorkTemp)
	}

	return
}

// ApiPoemSearch 古诗词搜索
func (ps *PoemService) ApiPoemSearch(tType, page int, value string) (workList []model.RXWorks, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	db := global.GVA_DB.Model(&model.RXWorks{}).Debug()
	switch tType {
	case 1:
		//古诗词标题
		db = db.Where("title like ?", "%"+value+"%")
	case 2:
		//古诗作者
		db = db.Where("author like ?", "%"+value+"%")
	case 3:
		//古诗朝代
		db = db.Where("dynasty = ?", value)
	case 4:
		//古诗类型
		db = db.Where("kind_cn = ?", value)
	default:
		//古诗词标题
		db = db.Where("title like ?", "%"+value+"%")
	}
	db = db.Order("id asc")
	db = db.Count(&total)
	db = db.Limit(size).Offset(offset)
	db.Find(&workList)
	for idx, _ := range workList {
		// 找到子串substr在str中的位置
		index := strings.Index(workList[idx].Content, "。")
		if index != -1 {
			// 使用切片操作截取字符串直到substr的结束位置
			workList[idx].ShortContent = workList[idx].Content[:index+len("。")]
		} else {
			workList[idx].ShortContent = workList[idx].Content
		}
	}
	return
}

// ApiPoemSearchList 古诗词搜索-多条件
func (ps *PoemService) ApiPoemSearchList(title, author, dynasty, kindCn string, page int) (workList []model.RXWorks, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	db := global.GVA_DB.Model(&model.RXWorks{}).Debug()

	if title != "" {
		db = db.Where("title like ?", "%"+title+"%")
	}
	if author != "" {
		db = db.Where("author like ?", "%"+author+"%")
	}
	if kindCn != "" {
		db = db.Where("kind_cn like ?", "%"+kindCn+"%")
	}
	if dynasty != "" {
		db = db.Where("dynasty = ?", "%"+dynasty+"%")
	}
	db = db.Order("id asc")
	db = db.Count(&total)
	db = db.Limit(size).Offset(offset)
	db.Find(&workList)

	for idx, _ := range workList {
		// 找到子串substr在str中的位置
		index := strings.Index(workList[idx].Content, "。")
		if index != -1 {
			// 使用切片操作截取字符串直到substr的结束位置
			workList[idx].ShortContent = workList[idx].Content[:index+len("。")]
		} else {
			workList[idx].ShortContent = workList[idx].Content
		}
	}
	return
}

// ApiPoemInfo 获取古诗词详情
func (ps *PoemService) ApiPoemInfo(workId int) (data model.RXWorks) {
	db := global.GVA_DB.Model(&model.RXWorks{}).Debug().Where("work_id = ?", workId)
	db.First(&data)
	return
}

// ApiAuthorList 获取作者列表
func (ps *PoemService) ApiAuthorList(dynasty string, page int) (authorList []model.RXAuthors, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	db := global.GVA_DB.Model(&model.RXAuthors{}).Debug().Where("dynasty = ?", dynasty)
	db = db.Order("id asc").Count(&total)
	db = db.Limit(size).Offset(offset)
	db.Find(&authorList)

	return
}

// ApiAuthorInfo 获取作者详情
func (ps *PoemService) ApiAuthorInfo(authorId int) (data model.RXAuthors) {
	db := global.GVA_DB.Model(&model.RXAuthors{}).Debug().Where("author_id = ?", authorId)
	db.First(&data)
	return
}

type SayingResponse struct {
	QuoteId     int    `json:"quote_id"`
	Quote       string `json:"quote"`
	QuoteAuthor string `json:"quote_author"`
	QuoteWork   string `json:"quote_work"`
	QuoteWorkId int    `json:"quote_work_id"`
	Collection  string `json:"collection"`
	Dynasty     string `json:"dynasty"`
}

// ApiSayingList 名言警句
func (ps *PoemService) ApiSayingList(dynasty, author string, page int) (sayingList []SayingResponse, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)

	var RXCollectionQuoteList []model.RXCollectionQuotes
	db := global.GVA_DB.Model(&model.RXCollectionQuotes{}).Debug()
	if dynasty != "" && author != "" {
		db.Joins("LEFT JOIN rx_quotes ON rx_collection_quotes.quote_id = rx_quotes.quote_id AND rx_quotes.dynasty = ? AND rx_quotes.author = ?", dynasty, author)
	}
	if dynasty != "" {
		db.Joins("LEFT JOIN rx_quotes ON rx_collection_quotes.quote_id = rx_quotes.quote_id AND rx_quotes.dynasty = ?", dynasty)
	}
	if author != "" {
		db.Joins("LEFT JOIN rx_quotes ON rx_collection_quotes.quote_id = rx_quotes.quote_id AND rx_quotes.author = ?", author)
	}

	db = db.Order("id asc").Count(&total)
	db = db.Limit(size).Offset(offset)
	db.Find(&RXCollectionQuoteList)

	var quoteIds, collectionIds []any
	for idx, _ := range RXCollectionQuoteList {
		quoteIds = append(quoteIds, RXCollectionQuoteList[idx].QuoteId)
		collectionIds = append(collectionIds, RXCollectionQuoteList[idx].CollectionId)
	}
	var RXQuotesList []model.RXQuotes
	global.GVA_DB.Model(&model.RXQuotes{}).Debug().Where("quote_id in(?)", utils.SliceUnique(quoteIds)).Find(&RXQuotesList)

	var RXCollectionList []model.RXCollections
	global.GVA_DB.Model(&model.RXCollections{}).Debug().Where("collection_id in(?)", utils.SliceUnique(collectionIds)).Find(&RXCollectionList)

	var sayingTemp SayingResponse
	for idx, _ := range RXCollectionQuoteList {
		sayingTemp.QuoteId = RXCollectionQuoteList[idx].QuoteId
		sayingTemp.Quote = RXCollectionQuoteList[idx].Quote
		sayingTemp.QuoteAuthor = RXCollectionQuoteList[idx].QuoteAuthor
		sayingTemp.QuoteWork = RXCollectionQuoteList[idx].QuoteWork
		sayingTemp.QuoteWorkId = RXCollectionQuoteList[idx].QuoteWorkId

		for idIndex, _ := range RXQuotesList {
			if RXQuotesList[idIndex].QuoteId == RXCollectionQuoteList[idx].QuoteId {
				sayingTemp.Dynasty = RXQuotesList[idIndex].Dynasty
			}
		}

		for idIndex, _ := range RXCollectionList {
			if RXCollectionList[idIndex].CollectionId == RXCollectionQuoteList[idx].CollectionId {
				sayingTemp.Collection = RXCollectionList[idIndex].Name
			}
		}

		sayingList = append(sayingList, sayingTemp)
	}
	return
}
