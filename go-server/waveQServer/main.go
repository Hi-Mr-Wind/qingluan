package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"waveQServer/src/comm"
	"waveQServer/src/config"
	"waveQServer/src/controller"
	"waveQServer/src/utils/logutil"
)

func main() {
	logutil.LogInfo("QingLuan is starting.....")
	logutil.LogInfo(comm.BANNER)
	logutil.LogInfo("version: V" + comm.VERSION)
	args := os.Args
	if len(args) == 1 {
		config.ReadConfiguration("")
	} else {
		config.ReadConfiguration(args[1])
	}
	gin.SetMode(gin.ReleaseMode) //开启生产环境
	logutil.LogInfo("QingLuan is started successfully. Port number:" + strconv.Itoa(int(config.GetConfig().Port)))
	//启动gin服务
	controller.Start(":" + strconv.Itoa(int(config.GetConfig().Port)))
}
