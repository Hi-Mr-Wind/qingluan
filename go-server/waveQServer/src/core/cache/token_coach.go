package cache

import (
	"fmt"
	"strings"
	"sync"
	"waveQServer/src/core/database"
	"waveQServer/src/entity"
)

type TokenPermissionsCache struct {
	cache map[string][]string
	mutex sync.Mutex
}

var (
	instance *TokenPermissionsCache
	once     sync.Once
)

// GetTokenInstance 返回单例模式的TokenPermissionsCache实例
func GetTokenInstance() *TokenPermissionsCache {
	once.Do(func() {
		instance = &TokenPermissionsCache{
			cache: make(map[string][]string),
		}
	})
	return instance
}

func NewTokenPermissionsCache() *TokenPermissionsCache {
	return &TokenPermissionsCache{
		cache: make(map[string][]string),
	}
}

// AddToken 添加新的令牌
func (c *TokenPermissionsCache) AddToken(token *entity.TokenPermission, permissions []string) {
	c.cache[token.Token] = permissions
}

func (c *TokenPermissionsCache) LoadFromDatabase() error {
	// 查询数据库获取令牌和对应的权限
	var tokens []struct {
		Token      string
		Permission string
	}
	if err := database.GetDb().Table("token_permissions").Select("token, permission").Find(&tokens).Error; err != nil {
		return fmt.Errorf("failed to load token permissions from database: %w", err)
	}

	// 遍历查询结果，将令牌和权限缓存到内存中
	for _, t := range tokens {
		permissions := strings.Split(t.Permission, ",") // 字符串转化为列表格式
		c.cache[t.Token] = permissions
	}
	return nil
}

// 获取token中的权限
func (c *TokenPermissionsCache) GetPermission(token string) ([]string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	permission, ok := c.cache[token]
	return permission, ok
}

// 获取token中的权限
func (c *TokenPermissionsCache) GetTokenPermission(tokens []string, permission string) bool {
	for i := range tokens {
		if string(i) == permission {
			return true
		}
	}
	return false
}

// DeleteToken 删除指定的令牌
func (c *TokenPermissionsCache) DeleteToken(token string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.cache, token)
	err := database.GetDb().Table("token_permissions").Delete("token = ?", token).Error
	return err
}

// TokenExists 检查指定的令牌是否存在
func (c *TokenPermissionsCache) TokenExists(token string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, exists := c.cache[token]
	return exists
}

// todo 需要自动清楚过期的令牌，但是过期的令牌并没有设置过期时间
// deleteExpiredTokens 方法

func init() {
	// 创建令牌权限缓存实例
	cache := NewTokenPermissionsCache()
	// 从数据库加载令牌权限到缓存中
	if err := cache.LoadFromDatabase(); err != nil {
		panic(err)
	}
}
