package utils

import (
	"fmt"
	"time"
)

const (
	timeFormatTpl = "2006-01-02"
	tf2           = "2006-01-02 15:04:05"
	RFC3339       = "2006-01-02T15:04:05Z07:00"
)

//FormatDateFromUnix 时间戳转格式化
func FormatDateFromUnix(unix int64) string {
	date := time.Unix(unix, 0)
	return date.Format(tf2)
}

//GetCurrentFormatDateTime 获取当前的日期时间
func GetCurrentFormatDateTime() string {
	return time.Now().Format(timeFormatTpl)
}

//GetCurrentFormatDate 获取当前的日期时间
func GetCurrentFormatDate() string {
	return time.Now().Format(tf2)
}

//MakeBetweenDateList 根据开始日期和结束日期计算出时间段内所有的日期，参数为日期格式：2023-02-09
func MakeBetweenDateList(sDate, eDate string) []string {
	var dateList []string

	if sDate == eDate {
		dateList = append(dateList, eDate)
		return dateList
	}
	date1, err := time.Parse(timeFormatTpl, sDate)
	if err != nil {
		fmt.Printf("时间解析异常,sDate(%v) err(%v)", sDate, err)
		return dateList
	}
	date2, err := time.Parse(timeFormatTpl, eDate)
	if err != nil {
		fmt.Printf("时间解析异常,eDate(%v) err(%v)", eDate, err)
		return dateList
	}
	if date2.Before(date1) {
		fmt.Printf("时间解析异常,sDate(%v) eDate(%v)", sDate, eDate)
		return dateList
	}
	// 输出日期格式固定
	date1Str := date1.Format(timeFormatTpl)
	date2Str := date2.Format(timeFormatTpl)
	dateList = append(dateList, date2Str)
	for {
		date2 = date2.AddDate(0, 0, -1)
		dateStr := date2.Format(timeFormatTpl)
		dateList = append(dateList, dateStr)
		if dateStr == date1Str {
			break
		}
	}
	return dateList
}
