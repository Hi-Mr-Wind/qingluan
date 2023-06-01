package main

import (
	"os"
	"strconv"
	"waveQServer/config"
	"waveQServer/router"
	"waveQServer/utils/logutil"
)

var strings = make(chan string)

func main() {
	logutil.LogInfo("QingLuan is starting.....")
	logutil.LogInfo(banner)
	logutil.LogInfo("version: V" + version)
	args := os.Args
	if len(args) == 1 {
		config.ReadConfiguration("")
	} else {
		config.ReadConfiguration(args[1])
	}
	//gin.SetMode(gin.ReleaseMode) //开启生产环境
	logutil.LogInfo("QingLuan is started successfully. Port number:" + strconv.Itoa(int(config.GetConfig().Port)))
	//启动gin服务
	router.Start(":" + strconv.Itoa(int(config.GetConfig().Port)))
}
