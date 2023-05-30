package main

import (
	"os"
	"strconv"
	"waveQServer/config"
	"waveQServer/core/service"
	"waveQServer/utils/logutil"
)

var strings = make(chan string)

func main() {
	logutil.LogInfo("QingLuan is starting.....")
	args := os.Args
	if len(args) == 1 {
		config.ReadConfiguration("")
	} else {
		config.ReadConfiguration(args[1])
	}
	//gin.SetMode(gin.ReleaseMode) //开启生产环境
	logutil.LogInfo("QingLuan is started successfully. Port number:" + strconv.Itoa(int(config.GetConfig().Port)))
	err := service.GetServer().Run(":" + strconv.Itoa(int(config.GetConfig().Port)))
	if err != nil {
		logutil.LogError("Service startup failure:" + err.Error())
		return
	}
	//
	//fmt.Println(banner)
	//fmt.Println(version)
	//gro, err := queue.NewGroup([]byte("qweasd"))
	//if err != nil {
	//	return
	//}
	//_, err = queue.New([]byte("qweasd"), []byte("123123123"))
	//if err != nil {
	//	return
	//}
	//
	//id, err := gro.GetGroupQueueById([]byte("123123123"))
	//if err != nil {
	//	return
	//}
	//go mes(id)
	//go get(id)
	////fmt.Println(<-strings)
	//fmt.Scanln()
}
