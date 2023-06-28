package utils

import (
	"runtime"
	"sync"
	"time"
)

var cac *Cache = NewCache()

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

type Cache struct {
	items map[string]cacheItem
	mutex sync.RWMutex
	stop  chan bool
}

// NewCache 创建一个新的缓存实例
func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]cacheItem),
		stop:  make(chan bool),
	}
	// 启动一个周期性的后台清理过期键的 goroutine
	go c.startCleanup()

	return c
}

// GetCache 获取一个预设的单例缓存实例
func GetCache() *Cache {
	return cac
}

// Get 获取缓存值
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	//惰性删除，每次查询到过期的key时候则触发删除
	if item.expiration.Before(time.Now()) {
		c.deleteInternal(key)
		return nil, false
	}

	return item.value, true
}

// Set 添加缓存值
// duration -缓存有效期  小于0则为永久
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var expiration time.Time
	if duration > 0 {
		expiration = time.Now().Add(duration)
	}

	c.items[key] = cacheItem{
		value:      value,
		expiration: expiration,
	}
}

// HasSet 判断key是否存在，这东西并不是一定靠谱
// 他返回的true不一定存在，但是返回false是绝对不存在
func (c *Cache) HasSet(key string) bool {
	_, ok := c.items[key]
	return ok
}

// Delete 删除缓存中的键值对
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.deleteInternal(key)
}

func (c *Cache) deleteInternal(key string) {
	delete(c.items, key)
}

// Clear 清空缓存
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items = make(map[string]cacheItem)
	runtime.GC() //通知gc回收
}

// 周期性清理过期键
func (c *Cache) startCleanup() {
	ticker := time.NewTicker(time.Minute) // 每分钟触发一次清理

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()

			c.mutex.Lock()

			for key, item := range c.items {
				if item.expiration.Before(currentTime) {
					c.deleteInternal(key)
				}
			}

			c.mutex.Unlock()
		case <-c.stop:
			ticker.Stop()
			return
		}
	}
}

// StopCleanup 停止过期清理
func (c *Cache) StopCleanup() {
	c.stop <- true
}
