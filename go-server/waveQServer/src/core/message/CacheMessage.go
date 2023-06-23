package message

import (
	"sync"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
	"waveQServer/src/utils"
)

// SetCachedSubMessage 创建一个订阅消息缓存
// 消息缓存虽说实现了同一个接口，但是内部属性并不完全一致，如果使用反射获取接口的类型则会损失一定的性能，所以这里选择多个方法分别进行缓存
func SetCachedSubMessage(mes *SubMessage) {
	mess := dto.SubMessage{
		MessageId:  mes.MessageId,
		ProducerId: mes.ProducerId,
		Timestamp:  mes.Timestamp,
		SendTime:   mes.SendTime,
		QueueId:    mes.QueueId,
		Indate:     mes.Indate,
		Subscriber: utils.ListToStr(mes.Subscriber),
		Body:       mes.Body,
	}
	database.GetDb().Create(&mess)
}

// SetCachedRandomMessage 创建一个权重随机消息的缓存
func SetCachedRandomMessage(mes *RandomMessage) {
	mess := dto.RandomMessage{
		MessageId:  mes.MessageId,
		ProducerId: mes.ProducerId,
		Timestamp:  mes.Timestamp,
		SendTime:   mes.SendTime,
		QueueId:    mes.QueueId,
		Indate:     mes.Indate,
		Weight:     mes.Weight,
		Number:     mes.Number,
		Body:       mes.Body,
	}
	database.GetDb().Create(&mess)
}

// SetCachedExclusiveMessage 创建一个独享消息的缓存队列
func SetCachedExclusiveMessage(mes *ExclusiveMessage) {
	mess := dto.ExclusiveMessage{
		MessageId:  mes.MessageId,
		ProducerId: mes.ProducerId,
		Timestamp:  mes.Timestamp,
		SendTime:   mes.SendTime,
		QueueId:    mes.QueueId,
		Indate:     mes.Indate,
		Body:       mes.Body,
	}
	database.GetDb().Create(&mess)
}

// SetCachedWeightMessage 创建一个权重消息缓存
func SetCachedWeightMessage(mes *WeightMessage) {
	mess := dto.WeightMessage{
		MessageId:  mes.MessageId,
		ProducerId: mes.ProducerId,
		Timestamp:  mes.Timestamp,
		SendTime:   mes.SendTime,
		QueueId:    mes.QueueId,
		Indate:     mes.Indate,
		Weight:     mes.Weight,
		Body:       mes.Body,
	}
	database.GetDb().Create(&mess)
}

// SetCachedDelayedMessage 创建一个延迟消息缓存
func SetCachedDelayedMessage(mes *DelayedMessage) {
	mess := dto.DelayedMessage{
		MessageId:  mes.MessageId,
		ProducerId: mes.ProducerId,
		Timestamp:  mes.Timestamp,
		SendTime:   mes.SendTime,
		QueueId:    mes.QueueId,
		Indate:     mes.Indate,
		Delayed:    mes.Delayed,
		Body:       mes.Body,
	}
	database.GetDb().Create(&mess)
}

// GetCachedSubMessage 根据消息ID获取一个订阅消息缓存
func GetCachedSubMessage(mesId string) SubMessage {
	messages := new(dto.SubMessage)
	database.GetDb().Find(messages, "message_id = ?", mesId).Scan(messages)
	return SubMessage{
		lock: sync.Mutex{},
		Heard: Heard{
			MessageId:  messages.MessageId,
			ProducerId: messages.ProducerId,
			QueueId:    messages.QueueId,
			Timestamp:  messages.Timestamp,
			SendTime:   messages.SendTime,
			Indate:     messages.Indate,
		},
		Subscriber: utils.StrToList(messages.Subscriber),
		Body:       messages.Body,
	}
}
