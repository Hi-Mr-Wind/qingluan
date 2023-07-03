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

// 消息状态
const (
	// NORMAL 正常
	NORMAL = iota
	// OVER_TIME 超时
	OVER_TIME
)

// 令牌权限
const (
	PermissionSubmitInfo  = "submit_info"
	PermissionCreateAPI   = "create_api"
	PermissionCreateGroup = "create_group"
	PermissionCreateQueue = "create_queue"
)
