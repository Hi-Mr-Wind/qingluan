package test

import (
	"fmt"
	"testing"
	"time"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
	"waveQServer/src/core/message"
	"waveQServer/src/utils"
	"waveQServer/src/utils/logutil"
)

func TestMd5(t *testing.T) {
	md5 := utils.Md5([]byte("Admin"))
	fmt.Println(md5)
	t.Log()
}

// 测试数据库连接
func TestDb(t *testing.T) {
	admin := new(dto.Admin)
	database.GetDb().Model(admin).Find(&admin)
	fmt.Println(*admin)
	admin.Id = "123123"
	fmt.Println(admin)
}

func TestLog(t *testing.T) {
	ch := make(chan int)
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			logutil.LogInfo("当前线程1：%v", time.Now().UnixNano())
		}
		ch <- 1
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			logutil.LogWarning("当前线程2：%v", time.Now().UnixNano())
		}
		ch1 <- 1
	}()
	<-ch
	<-ch1
}

func TestMes(t *testing.T) {
	//mes := message.SubMessage{
	//	Heard: message.Heard{
	//		MessageId:  "hefgbdfage",
	//		ProducerId: "asfqdacxasdqwcqwd",
	//		QueueId:    "c196de8f-d2d6-60d9-6092-c20433d156b3",
	//		Timestamp:  1687276300870,
	//		SendTime:   1687276300870,
	//		Indate:     0,
	//	},
	//	Subscriber: []string{"45a64421-35cc-6465-2046-3df88078542b"},
	//	Body:       []byte{2, 122, 98, 254},
	//}
	//message.SetCachedSubMessage(&mes)

	subMessage := message.GetCachedSubMessage("hefgbdfage")
	fmt.Println("查询到的数据为", subMessage)
}
