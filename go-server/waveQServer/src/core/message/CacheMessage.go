package message

import (
	"sync"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
	"waveQServer/src/utils"
)

// SetCachedSubMessage 创建一个订阅消息缓存
func SetCachedSubMessage(mes *Message) {
	message := dto.SubMessage{
		MessageId:  Heard{}.MessageId,
		ProducerId: Heard{}.ProducerId,
		Timestamp:  Heard{}.Timestamp,
		SendTime:   Heard{}.SendTime,
		QueueId:    Heard{}.QueueId,
		Indate:     Heard{}.Indate,
		Subscriber: utils.ListToStr(SubMessage{}.Subscriber),
		Body:       SubMessage{}.Body,
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
