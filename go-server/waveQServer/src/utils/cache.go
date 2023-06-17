package utils

import (
	"fmt"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
)

// 缓存用户apiKey对应的权限
var limits = make(map[string][]string)

// 缓存所有的apikey
var apiKeys []dto.User

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
func GetUser(apiKey string) *dto.User {
	for _, o := range apiKeys {
		if o.ApiKey == apiKey {
			return &o
		}
	}
	return nil
}

func init() {
	go func() {
		//查询所有的apikey并缓存
		user := new([]dto.User)
		database.GetDb().Find(&user)
		apiKeys = *user

		fmt.Println("加载apikey缓存完成", apiKeys)
		//加载apikey下的权限
		for _, o := range apiKeys {
			limit := new([]dto.QueueUeser)
			list := make([]string, 10)
			database.GetDb().Find(&limit, "user_id = ?", o.ApiKey)
			for _, v := range *limit {
				list = append(list, v.QueueId)
			}
			limits[o.ApiKey] = list
		}
		fmt.Println("加载权限缓存完成", limits)
	}()
}
