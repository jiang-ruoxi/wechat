package service

import (
	"sort"
	"wechat/model"
	"wechat/pkg/mysql"
)

type ChineseBookService struct {
}

type ChineseBookInfo struct {
	BookId    int    `json:"book_id"`
	BookCount string `json:"book_count"`
}

type ChineseBookInfoList struct {
	Id        int    `json:"id"`
	BookId    int    `json:"book_id"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Level     uint8  `json:"level"`
	Position  uint8  `json:"position"`
	BookCount string `json:"book_count"`
}

// GetChineseBookList 绘本列表
func (cbs *ChineseBookService) GetChineseBookList(level, page, pageSize int) (bookInfoList []ChineseBookInfoList, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	// 创建db
	var bookList []model.ChineseBook
	db := mysql.DB.Model(&model.ChineseBook{}).Debug()
	db = db.Where("type = ?", level)
	err = db.Count(&total).Error
	db = db.Order("position desc")
	db = db.Limit(limit).Offset(offset).Find(&bookList)

	db1 := mysql.DB.Model(&model.ChineseBookInfo{}).Debug()
	var bookDataList []ChineseBookInfo
	db1.Raw("SELECT book_id,count(id) as book_count FROM s_chinese_picture_info GROUP BY book_id").Scan(&bookDataList)

	var temp ChineseBookInfoList
	for _, item := range bookList {
		temp.Id = item.Id
		temp.BookId = item.BookId
		temp.Title = item.Title
		temp.Icon = item.Icon
		temp.Level = item.Level
		temp.Position = item.Position
		bookInfoList = append(bookInfoList, temp)
	}

	for index, item := range bookInfoList {
		for _, it := range bookDataList {
			if item.BookId == it.BookId {
				bookInfoList[index].BookCount = it.BookCount
			}
		}
	}

	sort.Slice(bookInfoList, func(i, j int) bool {
		if bookInfoList[i].Position > bookInfoList[j].Position {
			return true
		}
		return bookInfoList[i].Position == bookInfoList[j].Position && bookInfoList[i].Id < bookInfoList[j].Id
	})
	err = db.Error

	return bookInfoList, total, err
}

func (cbs *ChineseBookService) GetChineseBookInfo(bookId int) (list []model.ChineseBookInfo) {
	// 创建db
	var bookInfoList []model.ChineseBookInfo
	db := mysql.DB.Model(&model.ChineseBookInfo{}).Debug()
	db = db.Where("book_id = ?", bookId)
	db = db.Order("id asc")
	db = db.Find(&bookInfoList)
	return bookInfoList
}