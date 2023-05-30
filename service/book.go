package service

import (
	"wechat/model"
	"wechat/pkg/mysql"
)

type BookService struct {
}

// GetBookList 绘本列表
func (bs *BookService) GetBookList(level, page, pageSize int) (list []model.Book, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	// 创建db
	var bookList []model.Book
	db := mysql.DB.Model(&model.Book{}).Debug()
	db = db.Where("level = ?", level)
	err = db.Count(&total).Error
	db = db.Order("book_id asc")
	db = db.Limit(limit).Offset(offset).Find(&bookList)

	return bookList, total, err
}
