package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"waveQServer/config"
	"waveQServer/entity"
	"waveQServer/entity/queue"
	"waveQServer/utils/logutil"
)

var strings = make(chan string)

func main() {
	logutil.LogInfo("QingLuan is starting.....")
	args := os.Args
	_, err := config.ReadConfiguration(args[0])
	if err != nil {
		logutil.LogWarning(err.Error())
	}

	logutil.LogInfo("QingLuan is started successfully. Port number:" + strconv.Itoa(int(config.GetConfig().Port)))
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

func mes(id *queue.Queue) {
	for i := 1; i <= 100; i++ {
		heard := entity.NewHeard(nil, []byte("123123123"))
		message := entity.Message{
			Header: *heard,
			Body:   []byte("哈哈哈哈哈哈哈" + strconv.Itoa(i)),
		}
		id.Push(&message)
		fmt.Println("现存元素数量" + strconv.Itoa(int(id.Size())))
	}
}

func get(id *queue.Queue) {
	for {
		time.Sleep(10000)
		fmt.Println("剩余元素数量：" + strconv.Itoa(int(id.Size())))
		pull, err := id.Pull()
		if err != nil {
			fmt.Println(err.Error())
			continue
			//strings <- err.Error()
		}
		if pull.Body != nil {
			fmt.Println("消费到消息：" + string(pull.Body))
			//return strings
		}
		//return strings
	}
}
