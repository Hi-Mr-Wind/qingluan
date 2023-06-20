package enum

const (
	// OK 成功
	OK = 0
	// FAIL 失败
	FAIL = -1
)

const (
	// RandomMessage 随机消息
	RandomMessage = iota
	// SubMessage 订阅消息
	SubMessage
	// ExclusiveMessage 独享消息
	ExclusiveMessage
	// WeightMessage 权重消息
	WeightMessage
	// DelayedMessage 延迟消息
	DelayedMessage
)
