package service

import (
	"sort"
	"wechat/common/response"
	"wechat/global"
	"wechat/model"
)

type ChineseService struct {
}

// ApiBookNavList 获取国学绘本的列表信息
func (cs *ChineseService) ApiBookNavList(typeId int) (navList []model.ChineseBookNavName) {
	bookDB := global.GVA_DB.Model(&model.ChineseBookNavName{}).Debug()
	bookDB = bookDB.Where("status=1 and s_type = ?", typeId).Order("s_sort asc").Order("id asc")
	bookDB.Find(&navList)
	return navList
}

// GetChineseBookList 获取国学绘本的列表信息
func (cs *ChineseService) GetChineseBookList(level, page int) (chineseBookList []response.ResponseChineseBook, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	var bookList []model.ChineseBook
	bookDB := global.GVA_DB.Model(&model.ChineseBook{}).Debug()
	bookDB = bookDB.Where("type = ? and status = 1", level).Count(&total)
	bookDB = bookDB.Order("position desc").Limit(size).Offset(offset)
	bookDB.Find(&bookList)

	var bookInfoCountList []response.ResponseBookInfoCount
	infoDB := global.GVA_DB.Model(&model.ChineseBookInfo{}).Debug()
	infoDB.Raw("SELECT book_id,count(id) as book_count FROM s_chinese_picture_info GROUP BY book_id").Scan(&bookInfoCountList)

	var temp response.ResponseChineseBook
	for _, item := range bookList {
		temp.Id = item.Id
		temp.BookId = item.BookId
		temp.Title = item.Title
		temp.Icon = item.Icon
		temp.Level = item.Type
		temp.Position = item.Position
		chineseBookList = append(chineseBookList, temp)
	}

	for index, item := range chineseBookList {
		for _, it := range bookInfoCountList {
			if item.BookId == it.BookId {
				chineseBookList[index].BookCount = it.BookCount
			}
		}
	}

	sort.Slice(chineseBookList, func(i, j int) bool {
		if chineseBookList[i].Position > chineseBookList[j].Position {
			return true
		}
		return chineseBookList[i].Position == chineseBookList[j].Position && chineseBookList[i].Id < chineseBookList[j].Id
	})

	return
}

// GetChineseBookAlbumLList 获取国学绘本专辑的列表信息
func (cs *ChineseService) GetChineseBookAlbumLList(page int) (chineseBookAlbumList []model.ChineseBookAlbum, total int64) {
	size := global.DEFAULT_PAGE_SIZE
	offset := size * (page - 1)
	bookDB := global.GVA_DB.Model(&model.ChineseBookAlbum{}).Debug().Where("status=1")
	bookDB = bookDB.Count(&total)
	bookDB = bookDB.Order("position desc").Limit(size).Offset(offset)
	bookDB.Find(&chineseBookAlbumList)

	return
}

// GetChineseBookAlbumListInfo 获取国学绘本专辑对应列表信息
func (cs *ChineseService) GetChineseBookAlbumListInfo(bookId string) (chineseAlbumInfoList []model.ChineseAlbumInfo, total int64) {
	bookDB := global.GVA_DB.Model(&model.ChineseAlbumInfo{}).Debug()
	bookDB = bookDB.Where("book_id = ?", bookId).Count(&total)
	bookDB = bookDB.Order("position desc")
	bookDB.Find(&chineseAlbumInfoList)

	return
}

// GetChineseBookInfo 获取国学绘本的详情信息
func (cs *ChineseService) GetChineseBookInfo(bookId string) (bookInfoItems []model.ChineseBookInfo) {
	db := global.GVA_DB.Model(&model.ChineseBookInfo{}).Debug()
	db = db.Where("book_id = ?", bookId).Order("position asc").Find(&bookInfoItems)
	return
}

// GetChineseBookAlbumInfo 获取国学绘本专辑的详情信息
func (cs *ChineseService) GetChineseBookAlbumInfo(id int) (bookInfoItem model.ChineseAlbumInfo) {
	db := global.GVA_DB.Model(&model.ChineseAlbumInfo{}).Debug()
	db = db.Where("id = ?", id).First(&bookInfoItem)
	return
}
