package fileUtil

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"waveQServer/src/utils/logutil"
)

var pathsArray = make([]string, 1, 10)

// WriteFile 写入文件
func WriteFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		logutil.LogError(err.Error())
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logutil.LogError(err.Error())
		}
	}(file)
	_, err = io.WriteString(file, string(data))
	if err != nil {
		logutil.LogError(err.Error())
		return err
	}
	return nil
}

// ReadFile 读取文件
func ReadFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logutil.LogError(err.Error())
		return nil, err
	}
	return file, nil
}

// GetFileSize 获取文件大小（KB）单位 近似值
func GetFileSize(filePath string) (int32, error) {
	fileInfo, err := os.Stat(filePath)
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
