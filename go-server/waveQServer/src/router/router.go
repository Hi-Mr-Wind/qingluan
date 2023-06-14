package router

import (
	"waveQServer/src/api/admin"
	"waveQServer/src/utils/httpUtils"
	"waveQServer/src/utils/logutil"
)

// Start 启动服务
func Start(port string) {
	// 加载admin路由组
	admin.Include()

	//按照指定端口启动服务
	err := httpUtils.GetServer().Run(port)
	if err != nil {
		logutil.LogError("Service startup failure:" + err.Error())
		return
	}
}
