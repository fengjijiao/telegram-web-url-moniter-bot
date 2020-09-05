package database

import (
	"time"
)

// timeDate
// 返回时间
func timeDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// timestampNano
// 返回时间戳
func timestampNano() int {
	return int(time.Now().UnixNano() / 1e6)
}