package config

import (
	"encoding/json"
	"os"
)

// Configuration 配置文件实体
type Configuration struct {
	//端口号
	Port int32 `json:"port"`
	//工作模式 (true 为复制模式 false 为负载模式)
	PatternCopy bool `json:"pattern"`
}

// NewConfiguration 构造方法，用于构建Configuration实体并赋予默认值
func NewConfiguration() *Configuration {
	conf := new(Configuration)
	conf.Port = 9627
	conf.PatternCopy = true
	return conf
}

// ReadConfiguration 读取配置文件
func ReadConfiguration(filename string) (conf *Configuration, err error) {
	var configuration = NewConfiguration()
	file, err := os.Open(filename)
	if err != nil {
		return configuration, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	err = json.NewDecoder(file).Decode(&configuration)
	if err != nil {
		return conf, err
	}
	return configuration, nil
}
