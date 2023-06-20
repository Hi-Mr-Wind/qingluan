package lastingUtils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"waveQServer/src/utils/logutil"
)

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

// Decode 解码二进制数据
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
