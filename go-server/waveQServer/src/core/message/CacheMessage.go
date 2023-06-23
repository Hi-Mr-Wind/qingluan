package message

import (
	"sync"
	"waveQServer/src/core/database"
	"waveQServer/src/entity"
	"waveQServer/src/utils"
)

// SetCachedSubMessage 创建一个订阅消息缓存
// 消息缓存虽说实现了同一个接口，但是内部属性并不完全一致，如果使用反射获取接口的类型则会损失一定的性能，所以这里选择多个方法分别进行缓存
func SetCachedSubMessage(mes *SubMessage) {
	mess := entity.SubMessage{
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
	mess := entity.RandomMessage{
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
	mess := entity.ExclusiveMessage{
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
	mess := entity.WeightMessage{
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
	mess := entity.DelayedMessage{
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
	messages := new(entity.SubMessage)
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

// GetCachedRandomMessage 获根据消息ID获取随机权重消息
func GetCachedRandomMessage(mesId string) RandomMessage {
	messages := new(entity.RandomMessage)
	database.GetDb().Find(messages, "message_id = ?", mesId).Scan(messages)
	return RandomMessage{
		lock: sync.Mutex{},
		Heard: Heard{
			MessageId:  messages.MessageId,
			ProducerId: messages.ProducerId,
			QueueId:    messages.QueueId,
			Timestamp:  messages.Timestamp,
			SendTime:   messages.SendTime,
			Indate:     messages.Indate,
		},
		Weight: messages.Weight,
		Number: messages.Number,
		Body:   messages.Body,
	}
}

// GetCachedExclusiveMessage 根据ID获取一个独享消息
func GetCachedExclusiveMessage(mesId string) ExclusiveMessage {
	messages := new(entity.ExclusiveMessage)
	database.GetDb().Find(messages, "message_id = ?", mesId).Scan(messages)
	return ExclusiveMessage{
		lock: sync.Mutex{},
		Heard: Heard{
			MessageId:  messages.MessageId,
			ProducerId: messages.ProducerId,
			QueueId:    messages.QueueId,
			Timestamp:  messages.Timestamp,
			SendTime:   messages.SendTime,
			Indate:     messages.Indate,
		},
		Body: messages.Body,
	}
}

// GetCachedWeightMessage 根据ID获取一个权重消息
func GetCachedWeightMessage(mesId string) WeightMessage {
	messages := new(entity.WeightMessage)
	database.GetDb().Find(messages, "message_id = ?", mesId).Scan(messages)
	return WeightMessage{
		lock: sync.Mutex{},
		Heard: Heard{
			MessageId:  messages.MessageId,
			ProducerId: messages.ProducerId,
			QueueId:    messages.QueueId,
			Timestamp:  messages.Timestamp,
			SendTime:   messages.SendTime,
			Indate:     messages.Indate,
		},
		Weight: messages.Weight,
		Body:   messages.Body,
	}
}

// GetCachedDelayedMessage 根据ID获取一个延迟消息
func GetCachedDelayedMessage(mesId string) DelayedMessage {
	messages := new(entity.DelayedMessage)
	database.GetDb().Find(messages, "message_id = ?", mesId).Scan(messages)
	return DelayedMessage{
		lock: sync.Mutex{},
		Heard: Heard{
			MessageId:  messages.MessageId,
			ProducerId: messages.ProducerId,
			QueueId:    messages.QueueId,
			Timestamp:  messages.Timestamp,
			SendTime:   messages.SendTime,
			Indate:     messages.Indate,
		},
		Delayed: messages.Delayed,
		Body:    messages.Body,
	}
}
