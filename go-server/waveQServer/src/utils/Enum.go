package utils

// 队列类型枚举
const (
	// STANDARD 标准模式
	STANDARD = iota
	// DELAY 延迟模式
	DELAY
	// BLOCK 阻塞模式
	BLOCK
)

// 消息状态
const (
	// NORMAL 正常
	NORMAL = iota
	// OVER_TIME 超时
	OVER_TIME
)
