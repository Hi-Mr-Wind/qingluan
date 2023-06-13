package lastingUtils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"waveQServer/entity/message"
	"waveQServer/utils"
	"waveQServer/utils/logutil"
)

// AsyncMessage 异步将消息写入文件
func AsyncMessage(mes *message.Message) {
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
func GetMessageFromFile() ([]*message.Message, error) {
	mess := make([]*message.Message, 50)
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
		var msg message.Message
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

// CreateData 获取系统缓存的路径
func CreateData() *os.File {
	s := logutil.GetPath() + ".data"
	f, err := os.OpenFile(s, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

// Encode 数据编码为二进制
func Encode(a any) {
	buf := new(bytes.Buffer)
	// 创建一个 Gob 编码器
	enc := gob.NewEncoder(buf)
	err := enc.Encode(a)
	if err != nil {
		logutil.LogError(err.Error())
		return
	}
}

func Decode(data *bytes.Buffer, types any) {
	// 创建一个 Gob 解码器
	dec := gob.NewDecoder(data)
	err := dec.Decode(&types)
	if err != nil {
		logutil.LogError(err.Error())
		return
	}
}

// GetOsType 获取当前操作系统类型
func GetOsType() string {
	return runtime.GOOS
}

// GetDataSql 根据系统类型获取应使用的数据库链接
func GetDataSql() *gorm.DB {
	s := logutil.GetPath() + string(filepath.Separator) + getOsType()
	db, err := gorm.Open(sqlite.Open(s), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// 根据系统类型获取数据库文件
func getOsType() string {
	switch GetOsType() {
	case "windows":
		return "windata.dll"
	case "linux":
		return "linuxdata"
	case "mac":
		return "macdata"
	}
	return "linuxdata"
}
