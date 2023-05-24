package entity

import "time"

type Response struct {
	//消息ID
	MesId []byte
	//接收到消息的时间
	MesTime time.Time
	//消息发送状态
	MesState []byte
}
