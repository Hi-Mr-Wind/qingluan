package message

import (
	"sync"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
	"waveQServer/src/utils"
)

// SetCachedSubMessage 创建一个订阅消息缓存
func SetCachedSubMessage(mes *SubMessage) {
	message := dto.SubMessage{
		MessageId:  mes.MessageId,
		ProducerId: mes.ProducerId,
		Timestamp:  mes.Timestamp,
		SendTime:   mes.SendTime,
		QueueId:    mes.QueueId,
		Indate:     mes.Indate,
		Subscriber: utils.ListToStr(mes.Subscriber),
		Body:       mes.Body,
	}
	database.GetDb().Create(&message)
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
