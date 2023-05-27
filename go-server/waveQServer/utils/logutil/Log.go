package logutil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// LogInfo 记录info级别日志
func LogInfo(message string) {
	s := " INFO: " + message
	logMes := pieceLog(s)
	i := GetLogFileIo()
	defer func(i *os.File) {
		err := i.Close()
		if err != nil {

		}
	}(i)
	log.SetOutput(i)
	log.Println(s)
	fmt.Println(logMes)
}

// LogWarning 记录warning级别日志
func LogWarning(message string) {
	s := " WARNING: " + message
	logMes := pieceLog(s)
	i := GetLogFileIo()
	defer func(i *os.File) {
		err := i.Close()
		if err != nil {

		}
	}(i)
	log.SetOutput(i)
	log.Println(s)
	fmt.Println(logMes)
}

// LogError 记录error级别日志
func LogError(message string) {
	s := " ERROR: " + message
	logMes := pieceLog(s)
	i := GetLogFileIo()
	defer func(i *os.File) {
		err := i.Close()
		if err != nil {

		}
	}(i)
	log.SetOutput(i)
	log.Println(s)
	fmt.Println(logMes)
}

// 拼合日志信息
func pieceLog(message string) string {
	return time.Now().Format("2006-01-02 03:04:05") + message
}

// GetLogFileIo 获取日志文件的io
func GetLogFileIo() *os.File {
	f, err := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

// getLogFileName 获取日志文件名称,从绝对路径开始
func getLogFileName() string {
	pat := GetPath() + "logs"
	_, err := os.Stat(pat)
	if err != nil {
		if err := os.MkdirAll(pat, 0777); err != nil {
		}
	}
	return pat + string(filepath.Separator) + time.Now().Format("2006-01-02") + ".log"
}

// GetPath 获取项目绝对路径末尾带路径分隔符
func GetPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	absPath, err := filepath.Abs(wd)
	if err != nil {
		panic(err)
	}
	return absPath + string(filepath.Separator)
}
