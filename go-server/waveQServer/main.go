package main

import (
	"os"
	"strconv"
	"waveQServer/src/comm"
	"waveQServer/src/config"
	"waveQServer/src/controller"
	"waveQServer/src/utils/logutil"
)

func main() {
	comm.Play.Add(1)
	logutil.LogInfo("QingLuan is loading.....")
	logutil.LogInfo(comm.BANNER)
	logutil.LogInfo("version: V" + comm.VERSION)
	args := os.Args
	if len(args) == 1 {
		config.ReadConfiguration("")
	} else {
		config.ReadConfiguration(args[1])
	}
	comm.Play.Wait()
	//启动gin服务
	controller.Start(":" + strconv.Itoa(int(config.GetConfig().Port)))
}
