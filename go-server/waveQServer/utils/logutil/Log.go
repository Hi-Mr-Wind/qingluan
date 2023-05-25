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
	logMes := pieceLog("INFO: " + message)
	log.SetOutput(getLogFileIo())
	log.Println(logMes)
	fmt.Println(logMes)
}

// LogWarning 记录warning级别日志
func LogWarning(message string) {
	logMes := pieceLog("WARNING: " + message)
	log.SetOutput(getLogFileIo())
	log.Println(logMes)
	fmt.Println(logMes)
}

// LogError 记录error级别日志
func LogError(message string) {
	logMes := pieceLog("ERROR: " + message)
	log.SetOutput(getLogFileIo())
	log.Println(logMes)
	fmt.Println(logMes)
}

// 拼合日志信息
func pieceLog(message string) string {
	return time.Now().Format("2006-01-02 03:04:05") + message
}

// 获取日志文件的io
func getLogFileIo() *os.File {
	f, err := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("The logutil file close fail...")
		}
	}(f)
	return f
}

// getLogFileName 获取日志文件名称,从绝对路径开始
func getLogFileName() string {
	return GetPath() + "logs" + string(filepath.Separator) + time.Now().Format("2006-01-02") + ".logutil"
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
