package logutil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	logLock = sync.Mutex{}
	fileLog *os.File
	logName string
)

// 初始化日志
func logInit(prefix string) *log.Logger {
	name := getLogFileName()
	if name != logName {
		err := fileLog.Close()
		if err != nil {
			panic(err)
		}
		logName = name
		fileLog = GetLogFileIo()
	}
	return log.New(fileLog, prefix, log.Lmicroseconds|log.Ldate)
}

// LogInfo 记录info级别日志
func LogInfo(message string, v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	info := logInit("[INFO] ")
	prints(info, message, v)
}

// LogWarning 记录warning级别日志
func LogWarning(message string, v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	info := logInit("[WARNING] ")
	prints(info, message, v)
}

// LogError 记录error级别日志
func LogError(message string, v ...any) {
	logLock.Lock()
	defer logLock.Unlock()
	info := logInit("[ERROR] ")
	prints(info, message, v)
}

// 拼合日志信息
func pieceLog(message string) string {
	return time.Now().Format("2006-01-02 03:04:05") + message
}

// GetLogFileIo 获取日志文件的io
func GetLogFileIo() *os.File {
	f, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
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

func init() {
	//初始化日志名称
	logName = getLogFileName()
	// 初始化日志io
	fileLog = GetLogFileIo()
}

func prints(logs *log.Logger, message string, v []any) {
	if v == nil || len(v) == 0 {
		fmt.Printf(message)
		logs.Printf(message)
		fmt.Println()
	} else {
		fmt.Printf(message, v)
		logs.Printf(message, v)
		fmt.Println()
	}
}
