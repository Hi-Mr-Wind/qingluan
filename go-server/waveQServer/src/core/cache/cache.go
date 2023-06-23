package cache

import (
	"time"
	"waveQServer/src/core/database"
	"waveQServer/src/core/message"
	"waveQServer/src/entity"
	"waveQServer/src/utils/logutil"
)

// 缓存用户apiKey对应的权限
var limits = make(map[string][]string)

// 缓存所有的apikey
var apiKeys []entity.User

// GetLimit 根据apikey获取所拥有的权限
func GetLimit(apiKey string) []string {
	return limits[apiKey]
}

// IsValidLimit 判断api是否持有某队列的权限
func IsValidLimit(apiKey string, queueId string) bool {
	return InSlice(GetLimit(apiKey), queueId)
}

// InSlice 判断元素是否在list中存在
func InSlice(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// IsInKeys 判断某apiKey是否存在
func IsInKeys(apikey string) bool {
	list := make([]string, 10)
	for _, o := range apiKeys {
		list = append(list, o.ApiKey)
	}
	return InSlice(list, apikey)
}

// GetUser 获取apiKey对应的对象
func GetUser(apiKey string) *entity.User {
	for _, o := range apiKeys {
		if o.ApiKey == apiKey {
			return &o
		}
	}
	return nil
}

// DelApiKey 移除内存中的 apikey以及对应权限
func DelApiKey(apiKey string) {
	var index int = -1
	for i, v := range apiKeys {
		if v.ApiKey == apiKey {
			index = i
			break
		}
	}
	if index != -1 {
		apiKeys = append(apiKeys[:index], apiKeys[index+1:]...)
		delete(limits, apiKey)
		delApikey(apiKey)
	}
}

// 删除持久化中的apikey
func delApikey(apikey string) {
	database.GetDb().Where("api_key = ?", apikey).Delete(&entity.User{})
	database.GetDb().Where("user_id = ?", apikey).Delete(&entity.QueueUeser{})
}

// 删除过期的消息
func deleteMessage() {
	milli := time.Now().UnixMilli()
	database.GetDb().Where("indate <= ?", milli).Delete(&message.SubMessage{})
	database.GetDb().Where("indate <= ?", milli).Delete(&message.WeightMessage{})
	database.GetDb().Where("indate <= ?", milli).Delete(&message.ExclusiveMessage{})
	database.GetDb().Where("indate <= ?", milli).Delete(&message.DelayedMessage{})
	database.GetDb().Where("indate <= ?", milli).Delete(&message.RandomMessage{})
}

// 删除过期的apikey
func deletePastDueApiKey() {
	milli := time.Now().UnixMilli()
	user := new([]entity.User)
	database.GetDb().Where("expiration_time <= ?", milli).Find(user) //找到所有过期的apikey
	for _, v := range *user {
		delApikey(v.ApiKey)
	}
}

// 加载apikey缓存
func loadingCache() {
	//查询所有的apikey并缓存
	user := new([]entity.User)
	database.GetDb().Find(&user)
	apiKeys = *user
	//加载apikey下的权限
	for _, o := range apiKeys {
		limit := new([]entity.QueueUeser)
		list := make([]string, 10)
		database.GetDb().Find(&limit, "user_id = ?", o.ApiKey)
		for _, v := range *limit {
			list = append(list, v.QueueId)
		}
		limits[o.ApiKey] = list
	}
}

func init() {
	go func() {
		//删除未被清理的过期的apikey
		logutil.LogInfo("Start deleting expired apikey……")
		deletePastDueApiKey()
		logutil.LogInfo("Deleting expired apikey is complete！")
		//删除未被清理的过期消息
		logutil.LogInfo("Start deleting expired messages……")
		deleteMessage()
		logutil.LogInfo("Deleting expired messages is complete！")
		//加载apikey缓存
		logutil.LogInfo("Loading the apikey cache……")
		loadingCache()
		logutil.LogInfo("loading the apikey cache is complete!")
	}()
}
