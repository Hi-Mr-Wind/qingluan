package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
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
