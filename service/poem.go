package service

import (
	"wechat/global"
	"wechat/model"
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
	KindId int    `json:"kind_id"`
	Kind   string `json:"kind"`
}

// ApiKindList 古诗词集合类别列表
func (ps *PoemService) ApiKindList() (KindList []KindResponse) {
	var collectionModelList []model.RXCollections
	db := global.GVA_DB.Model(&model.RXCollections{}).Debug()
	db = db.Order("sort desc")
	db.Group("kind_id").Find(&collectionModelList)
	var kindTemp KindResponse
	for idx, _ := range collectionModelList {
		kindTemp.KindId = collectionModelList[idx].KindId
		kindTemp.Kind = collectionModelList[idx].Kind
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
