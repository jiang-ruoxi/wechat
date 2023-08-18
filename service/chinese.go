package service

import (
	"wechat/common/response"
	"wechat/global"
	"wechat/model"
	"sort"
)

type ChineseService struct {
}

//GetChineseBookList 获取国学绘本的列表信息
func (cs *ChineseService) GetChineseBookList(level, page, size int) (chineseBookList []response.ResponseChineseBook, total int64) {
	offset := size * (page - 1)
	var bookList []model.ChineseBook
	bookDB := global.GVA_DB.Model(&model.ChineseBook{}).Debug()
	bookDB = bookDB.Where("type = ?", level).Count(&total)
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

//GetChineseBookInfo 获取国学绘本的详情信息
func (cs *ChineseService) GetChineseBookInfo(bookId string) (bookInfoItems []model.ChineseBookInfo) {
	db := global.GVA_DB.Model(&model.ChineseBookInfo{}).Debug()
	db = db.Where("book_id = ?", bookId).Order("id asc").Find(&bookInfoItems)
	return
}
