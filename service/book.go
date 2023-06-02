package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"unsafe"
	"wechat/common"
	"wechat/model"
	"wechat/pkg/mysql"
	"wechat/pkg/redis"
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
	db = db.Order("position asc")
	db = db.Limit(limit).Offset(offset).Find(&bookList)

	return bookList, total, err
}
func (bs *BookService) GetBookInfo(bookId int) (list []model.BookInfo) {
	// 创建db
	var bookInfoList []model.BookInfo
	db := mysql.DB.Model(&model.BookInfo{}).Debug()
	db = db.Where("book_id = ?", bookId)
	db = db.Order("id asc")
	db = db.Find(&bookInfoList)
	return bookInfoList
}

func (bs *BookService) GetBookInfoBak(bookId int) {
	url := "https://kadacan-service.hhdd.com/book/getFullBookInfo"
	song := make(map[string]interface{})
	song["bookId"] = bookId
	song["userId"] = 16159273
	bytesData, err := json.Marshal(song)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("cookie", "_HHDD_CAN_=eyJleHRUaW1lIjoxNzE2OTc1NTg3NjMyLCJuaWNrIjoi6Zi/55m9IiwidG9rZW4iOiIwMjFhOTI1NDE5NDI0NTYwMDA0MTMwNzBmOGIwZTA5NiIsInVzZXJJZCI6MTYxNTkyNzN9; Max-Age=31536000; Expires=Wed, 29-May-2024 09:39:47 GMT; Path=/")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

	st := strconv.Itoa(bookId)

	fileName := st + ".log"

	// 打开文件并追加内容
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(*str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File content:", *str)
}

func (bs *BookService) InsertVideoLog(c *common.VideoLogReq) (err error) {
	//定义对应的类型
	var data model.VideoLog
	//格式化数据生成
	c.GenerateVideoLog(&data)
	if err = mysql.DB.Model(&model.VideoLog{}).Create(&data).Error; err != nil {
		fmt.Println("数据创建失败")
		return err
	}
	return nil
}

func (bs *BookService) MakeVideo(c *common.VideoLogReq) (err error) {
	//定义对应的类型
	var data model.VideoLog
	//格式化数据生成
	c.GenerateVideoLog(&data)

	var videoLog []model.VideoLog

	db := mysql.DB.Model(&model.VideoLog{}).Debug()
	db.Where("open_id =? and book_id = ?", data.OpenId, data.BookId).Find(&videoLog)

	if err = db.Error; err != nil {
		fmt.Println("查询失败")
		return err
	}

	videoLogJson, _ := json.Marshal(videoLog)
	item := string(videoLogJson)
	fmt.Printf("%#v \n",item)
	
	queue := "huiben_video_make"
	redis.RedisClient.RPush(context.Background(), queue, item).Result()

	return nil
}
