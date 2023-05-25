package api

import "waveQServer/entity"

// Query 队列接口
type Query interface {

	//获取消息
	getMessage() entity.Message

	//发送消息
	sendMessage(mes entity.Message) bool

	//发送延迟消息
	sendDelayMessage(mes entity.Message, time int32) bool
}
