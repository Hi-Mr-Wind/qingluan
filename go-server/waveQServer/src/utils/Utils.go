package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
	"waveQServer/src/utils/logutil"
)

type Def func()

// GetApiKey 根据权限和随机ID生成一个唯一性的apikey
func GetApiKey(recessRights []string) string {
	data := make([]byte, 50, 200)
	for i := 0; i < len(recessRights); i++ {
		for j := 0; j < len(recessRights[i]); j++ {
			data = append(data, recessRights[i][j])
		}
	}
	//拿到一个纳秒级别的16进制时间字符串，此处用纳秒是因为纳秒的时间精度极高，产生出重复的概率可以忽略不计
	id := []byte(strconv.FormatInt(time.Now().UnixNano(), 16))
	data = append(data, id...)
	// 创建一个 SHA256 的哈希实例
	hash := sha256.New()
	// 向哈希实例输入数据
	hash.Write(data)
	// 计算 SHA256 哈希值的字节数组
	hashBytes := hash.Sum(nil)
	//gc回收垃圾对象
	defer func() {
		data = nil
		runtime.GC()
	}()
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

// IsEmpty 字符串为空
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}

// IsNull 判断指针为nil
func IsNull(data *any) bool {
	return data == nil
}

// IsNotNull 判断指针为非nil
func IsNotNull(data *any) bool {
	return !IsNull(data)
}

// ToJsonString 结构体转json字符串
func ToJsonString(date any) string {
	data, err := json.Marshal(date)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return ""
	}
	return string(data)
}

// ListToStr 切片转字符串，以“,”分隔
func ListToStr(list []string) string {
	return strings.Join(list, ",")
}

// StrToList 将“，”分隔的字符串转化为切片
func StrToList(str string) []string {
	return strings.Split(str, ",")
}

// TimeTask 设置一个一次性的定时器
func TimeTask(t time.Duration, f Def) {
	time.AfterFunc(time.Millisecond*t, f)
}

// MapToJson map转为json
func MapToJson(m map[string]interface{}) string {
	// 将map转为JSON格式
	result, err := json.Marshal(m)
	if err != nil {
		logutil.LogError(err.Error())
		return ""
	}
	return string(result)
}

// JsonToMap json转map
func JsonToMap(j string) map[string]interface{} {
	data := map[string]interface{}{}
	// 解析JSON数据到map中
	if err := json.Unmarshal([]byte(j), &data); err != nil {
		logutil.LogError(err.Error())
		return nil
	}
	return data
}
