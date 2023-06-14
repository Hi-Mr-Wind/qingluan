package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"waveQServer/src/utils"
	"waveQServer/src/utils/logutil"
)

var config *Configuration

// Configuration 配置文件实体
type Configuration struct {
	//端口号
	Port int32 `json:"port"`
	//持久化目录
	DataPath string `json:"dataPath"`
	//默认持久化文件大小，超过此大小的的消息将被清除  单位MB
	DataFileSize int64 `json:"dataFileSize"`
	//心跳时间 单位秒
	HeartBeat int32 `json:"heartBeat"`
	//管理员用户名
	UserName string `json:"userName"`
	//管理员密码
	Password string `json:"password"`
	//消息持有时间，单位天
	PossessTime int32 `json:"possessTime"`
}

// NewConfiguration 构造方法，用于构建Configuration实体并赋予默认值
func NewConfiguration() *Configuration {
	conf := new(Configuration)
	conf.Port = 9627
	conf.DataPath = logutil.GetPath() + "data"
	conf.HeartBeat = 60
	conf.DataFileSize = 100 * 1024
	conf.UserName = utils.Md5([]byte("Admin"))
	conf.Password = utils.Md5([]byte("Admin"))
	config = conf
	return conf
}

// ReadConfiguration 读取配置文件
func ReadConfiguration(filename string) (conf *Configuration) {
	if filename == "" {
		filename = logutil.GetPath() + "config" + string(filepath.Separator) + "conf.json"
	}
	var configuration = NewConfiguration()
	if _, err := os.Stat(filename); err != nil {
		logutil.LogWarning("the configuration file does not exist. The default configuration has been enabled")
		return configuration
	}
	file, err := os.Open(filename)
	if err != nil {
		logutil.LogError(err.Error())
		return configuration
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logutil.LogError(err.Error())
		}
	}(file)

	err = json.NewDecoder(file).Decode(&configuration)
	if err != nil {
		return conf
	}
	if configuration.UserName == "" {
		configuration.UserName = utils.Md5([]byte("Admin"))
	} else {
		configuration.UserName = utils.Md5([]byte(configuration.UserName))
	}
	if configuration.Password == "" {
		configuration.Password = utils.Md5([]byte("Admin"))
	} else {
		configuration.Password = utils.Md5([]byte(configuration.Password))
	}
	config = configuration
	return configuration
}

// GetConfig 获取配置文件信息
func GetConfig() *Configuration {
	return config
}
