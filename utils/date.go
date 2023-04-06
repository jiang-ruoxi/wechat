package utils

import "time"

const (
	timeFormatTpl = "2006-01-02"
	tf2           = "2006-01-02 15:04:05"
	RFC3339       = "2006-01-02T15:04:05Z07:00"
)

// 时间戳转格式化
func FormatDateFromUnix(unix int64) string {
	date := time.Unix(unix, 0)
	return date.Format(tf2)
}
