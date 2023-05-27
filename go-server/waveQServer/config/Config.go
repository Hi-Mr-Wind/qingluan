package config

import (
	"encoding/json"
	"os"
	"waveQServer/utils/logutil"
)

var config *Configuration

// Configuration 配置文件实体
type Configuration struct {
	//端口号
	Port int32 `json:"port"`
	//工作模式 (true 为复制模式 false 为负载模式)
	PatternCopy bool `json:"pattern"`
	//持久化目录
	DataPath string `json:"dataPath"`
	//默认持久化文件大小，超过此大小的的消息将被清除  单位MB
	DataFileSize int64 `json:"dataFileSize"`
	//心跳时间 单位秒
	HeartBeat int32 `json:"heartBeat"`
}

// NewConfiguration 构造方法，用于构建Configuration实体并赋予默认值
func NewConfiguration() *Configuration {
	conf := new(Configuration)
	conf.Port = 9627
	conf.PatternCopy = true
	conf.DataPath = logutil.GetPath() + "data"
	conf.HeartBeat = 60
	conf.DataFileSize = 100 * 1024
	return conf
}

// ReadConfiguration 读取配置文件
func ReadConfiguration(filename string) (conf *Configuration, err error) {
	var configuration = NewConfiguration()
	file, err := os.Open(filename)
	if err != nil {
		logutil.LogError(err.Error())
		return configuration, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logutil.LogError(err.Error())
		}
	}(file)

	err = json.NewDecoder(file).Decode(&configuration)
	if err != nil {
		return conf, err
	}
	config = configuration
	return configuration, nil
}

// GetConfig 获取配置文件信息
func GetConfig() *Configuration {
	return config
}
