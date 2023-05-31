package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"waveQServer/utils/logutil"
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

// GetApiKey 根据权限和随机ID生成一个唯一性的apikey
func GetApiKey(rccessRights [][]byte) string {
	data := make([]byte, 50, 100)
	for i := 0; i < len(rccessRights); i++ {
		for j := 0; j < len(rccessRights[i]); j++ {
			data = append(data, rccessRights[i][j])
		}
	}
	id := []byte(uuid.New().String())
	data = append(data, id...)
	// 创建一个 SHA256 的哈希实例
	hash := sha256.New()
	// 向哈希实例输入数据
	hash.Write(data)
	// 计算 SHA256 哈希值的字节数组
	hashBytes := hash.Sum(nil)
	// 将字节数组格式化为十六进制字符串
	return fmt.Sprintf("%x", hashBytes)
}

// Md5 MD5加密
func Md5(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

// GetTime 获取一个格式化后的当前时间
func GetTime() string {
	return time.Now().Format("2006-01-02 03:04:05")
}

// Equals 判断两个可比较类型相等
func Equals[T comparable](t T, m T) bool {
	return t == m
}

// NotEquals  判断两个可比较类型不相等
func NotEquals[T comparable](t T, m T) bool {
	return !Equals(t, m)
}

func IsEmpty(data string) bool {
	if data == "" || len([]byte(data)) == 0 {
		return true
	}
	return false
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
