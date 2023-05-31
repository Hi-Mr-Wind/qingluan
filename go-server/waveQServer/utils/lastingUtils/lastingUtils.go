package lastingUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"waveQServer/entity"
	"waveQServer/utils"
	"waveQServer/utils/logutil"
)

// AsyncMessage 异步将消息写入文件
func AsyncMessage(mes *entity.Message) {
	marshal, err := json.Marshal(mes)
	if err != nil {
		logutil.LogError(err.Error())
		return
	}
	fileName := fmt.Sprintf("%sdata%s%s%s%s.data", logutil.GetPath(), string(filepath.Separator), string(mes.Header.QueueID), string(filepath.Separator), string(mes.Header.Id))
	err = utils.WriteFile(fileName, marshal)
	if err != nil {
		logutil.LogError(err.Error())
		return
	}
	mes.Header.File = fileName
}

// GetMessageFromFile 解析所有持久化的消息
func GetMessageFromFile() ([]*entity.Message, error) {
	mess := make([]*entity.Message, 50)
	// 读取文件内容
	files := utils.GetPathFiles(fmt.Sprintf("%sdata", logutil.GetPath()))
	if len(files) == 0 {
		return nil, errors.New("no cached message")
	}
	for _, filr := range files {
		msgBytes, err := ioutil.ReadFile(filr)
		if err != nil {
			return nil, err
		}
		// 解析为消息对象
		var msg entity.Message
		err = json.Unmarshal(msgBytes, &msg)
		if err != nil {
			return nil, err
		}
		mess = append(mess, &msg)
	}
	return mess, nil
}

// CreateQueueCatalogue 创建用于队列持久化的目录
func CreateQueueCatalogue(groupId string, queueId string) string {
	path := logutil.GetPath() + groupId + string(filepath.Separator) + queueId + string(filepath.Separator)
	err := os.MkdirAll(path, 0750)
	if err != nil {
		logutil.LogError(err.Error())
		return ""
	}
	return path
}
