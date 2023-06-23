package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"waveQServer/src/config"
	"waveQServer/src/controller/admin"
	"waveQServer/src/utils/httpUtils"
	"waveQServer/src/utils/logutil"
)

// Start 启动服务
func Start(port string) {
	gin.SetMode(gin.ReleaseMode) //开启生产环境
	// 加载admin路由组
	admin.Include()
	logutil.LogInfo("QingLuan is started successfully. Port number:" + strconv.Itoa(int(config.GetConfig().Port)))
	//按照指定端口启动服务
	err := httpUtils.GetServer().Run(port)
	if err != nil {
		logutil.LogError("Service startup failure:" + err.Error())
		return
	}
}
