package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"waveQServer/utils/logutil"
)

var pathsArray = make([]string, 1, 10)

// GetFileSize 获取文件大小（KB）单位 近似值
func GetFileSize(filePath string) (int32, error) {
	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		logutil.LogError(err.Error())
		return 0, err
	}
	size := fileInfo.Size()
	return int32(float64(size) / 1024), nil // 输出结果
}

// GetPathFiles 获取指定目录下所有.data文件
func GetPathFiles(path string) []string {
	err := filepath.Walk(path, walkFunc)
	if err != nil {
		logutil.LogError(err.Error())
		return nil
	}
	return pathsArray
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if !info.IsDir() && strings.HasSuffix(path, ".data") {
		pathsArray = append(pathsArray, path)
	}
	return nil
}
