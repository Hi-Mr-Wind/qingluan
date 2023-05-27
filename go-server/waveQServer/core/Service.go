package core

import (
	"github.com/gin-gonic/gin"
	"waveQServer/utils/logutil"
)

var server *gin.Engine

// GetServer 获取一个服务引擎
func GetServer() *gin.Engine {
	if server == nil {
		gin.DefaultWriter = logutil.GetLogFileIo() //输出日志切换为文件
		server = gin.New()
	}
	return server
}
