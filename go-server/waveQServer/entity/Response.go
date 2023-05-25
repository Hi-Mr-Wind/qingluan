package entity

import "time"

// Response 返回结构体
type Response struct {
	//消息ID
	MesId []byte
	//接收到消息的时间
	MesTime time.Time
	//消息发送状态
	MesState []byte
}
