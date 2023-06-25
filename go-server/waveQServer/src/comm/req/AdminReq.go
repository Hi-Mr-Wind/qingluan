package req

// CreateApiKeyReq 创建apikey接收参数
type CreateApiKeyReq struct {
	//权限
	RecessRights []string
	//过期时间
	ExpirationTime int64
}
