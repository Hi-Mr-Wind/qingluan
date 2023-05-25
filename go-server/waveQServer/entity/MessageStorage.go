package entity

import "time"

// Storage 持久化结构
type Storage struct {
	//消息内容
	Messages []*Message `json:"messages"`
	//持久化时间
	StorageTime time.Time `json:"storageTime"`
}
