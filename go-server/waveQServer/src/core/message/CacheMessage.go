package message

import (
	"reflect"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
	"waveQServer/src/utils"
)

// SetCachedMessage 创建一个消息缓存
func SetCachedMessage(mes *Message) {
	message := *mes
	header := message.GetHeader()
	messages := dto.Messages{
		MessageId: header.Id,
		DataJson:  utils.ToJsonString(message),
		DateType:  reflect.TypeOf(message).String(),
		QueueId:   header.QueueID,
		CreatTime: utils.GetTime(),
	}
	database.GetDb().Create(&messages)
}

func GetCachedMessage(mesId string) *Message {
	messages := new(dto.Messages)
	database.GetDb().Find(messages)

	return messages
}
