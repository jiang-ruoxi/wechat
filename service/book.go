package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
	"unsafe"
	"wechat/common"
	"wechat/model"
	"wechat/pkg/mysql"
	"wechat/pkg/redis"
)

type BookService struct {
}

type BookInfo struct {
	BookId    int    `json:"book_id"`
	BookCount string `json:"book_count"`
}

type BookInfoList struct {
	Id        int    `json:"id"`
	BookId    int    `json:"book_id"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Level     uint8  `json:"level"`
	Position  uint8  `json:"position"`
	BookCount string `json:"book_count"`
}

// GetBookList 绘本列表
func (bs *BookService) GetBookList(level, page, pageSize int) (bookInfoList []BookInfoList, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	// 创建db
	var bookList []model.Book
	db := mysql.DB.Model(&model.Book{}).Debug()
	db = db.Where("level = ?", level)
	err = db.Count(&total).Error
	db = db.Order("position desc")
	db = db.Limit(limit).Offset(offset).Find(&bookList)

	db1 := mysql.DB.Model(&model.BookInfo{}).Debug()
	var bookDataList []BookInfo
	db1.Raw("SELECT book_id,count(id) as book_count FROM s_huiben_info GROUP BY book_id").Scan(&bookDataList)

	var temp BookInfoList
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
	fmt.Printf("%#v \n", item)

	queue := "huiben_video_make"
	redis.RedisClient.RPush(context.Background(), queue, item).Result()

	return nil
}

//IsCheckSign 验证请求
func (bs *BookService) IsCheckSign(sign string) bool {
	return true
	if sign == "" || len(sign) != 32 {
		return false
	}
	t := time.Now().Unix()
	fmt.Println(t)
	var slice []int64 = make([]int64, 0, 10)
	for i := 0; i < 10; i++ {
		t = t - 1
		slice = append(slice, t)
	}

	md5StrSlice := make([]string, 0, 15)
	for _, item := range slice {
		str := strconv.FormatInt(item, 10)
		md := fmt.Sprintf("%s%s", str, "ruoxi")
		fmt.Println(md)
		hash := md5.Sum([]byte(md))
		md5str := hex.EncodeToString(hash[:])
		fmt.Println(md5str + "\n")
		md5StrSlice = append(md5StrSlice, md5str)
	}
	isExit := bs.IsInSlice(sign, md5StrSlice)
	if !isExit {
		return false
	}

	return true
}

//IsInSlice 判断字符串是否在切片中
func (bs *BookService) IsInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

type MiniData struct {
	Name    string `json:"name"`
	AppId   string `json:"app_id"`
	Icon    string `json:"icon"`
	Content string `json:"content"`
}

func (bs *BookService) GetMiniList() []MiniData {
	gj := MiniData{
		Name:    "百科知识公基",
		AppId:   "wxb1e2d179618ef271",
		Icon:    "https://static.58haha.com/mini/wxb1e2d179618ef271.jpeg",
		Content: "一个事业编刷题和事业编答题一站式的事业编题库百科知识增长，应对事业编刷题题库和事业编答题题库的小程序",
	}
	gx := MiniData{
		Name:    "国学拼音识数运算",
		AppId:   "wx1090960cb624449c",
		Icon:    "https://static.58haha.com/mini/wx1090960cb624449c.jpeg",
		Content: "一个事业编刷题和事业编答题一站式的事业编题库百科知识增长，应对事业编刷题题库和事业编答题题库的小程序",
	}
	fj := MiniData{
		Name:    "英语绘本分级跟读",
		AppId:   "wx65b5468d031d0923",
		Icon:    "https://static.58haha.com/mini/wx65b5468d031d0923.jpeg",
		Content: "一款英语分级绘本跟读小程序，可以实现英语绘本跟读，也可以进行英语绘本阅读的小程序",
	}
	cj := MiniData{
		Name:    "初级英语绘本跟读",
		AppId:   "wxd1adfdd4aa40cf1a",
		Icon:    "https://static.58haha.com/mini/wxd1adfdd4aa40cf1a.jpeg",
		Content: "一款儿童初级英语绘本跟读小程序，可以实现英语绘本跟读，也可以进行英语绘本阅读的小程序",
	}
	var listMiniData []MiniData
	listMiniData = append(listMiniData, gj, gx, fj, cj)

	return listMiniData
}
