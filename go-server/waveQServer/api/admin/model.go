package admin

import (
	"waveQServer/utils/httpUtils"
)

// Include admin路由组模块
func Include() {
	Enter(httpUtils.GetRouterGroup("/"))
	// admin路由组
	group := httpUtils.GetRouterGroup("/admin")
	group.Use(httpUtils.Token)
	Urls(group)
}
