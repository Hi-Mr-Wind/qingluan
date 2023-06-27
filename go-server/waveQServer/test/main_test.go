package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
	"waveQServer/src/comm"
	"waveQServer/src/comm/enum"
	"waveQServer/src/core/cache"
	"waveQServer/src/core/database"
	"waveQServer/src/core/message"
	"waveQServer/src/entity"
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
	admin := new(entity.Admin)
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

func TestApikey(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("当前内存占用量：%d KB\n", m.Alloc/1024)
	a := []string{"c196de8f-d2d6-60d9-6092-c20433d156b3", "12D1E", "cqw12的饿", "dqw3dasdqwd"}
	key := utils.GetApiKey(a)
	fmt.Println(key)

	// 再次读取内存占用情况
	runtime.ReadMemStats(&m)
	fmt.Printf("回收后内存占用量：%d KB\n", m.Alloc/1024)

	b := []string{"c196de8f-d2d6-60d9-6092-c20433d156b3", "test_group"}
	key = utils.GetApiKey(b)
	fmt.Println(key)

	// 再次读取内存占用情况
	runtime.ReadMemStats(&m)
	fmt.Printf("回收后内存占用量：%d KB\n", m.Alloc/1024)
}

func TestMes(t *testing.T) {
	cache.DelApiKey("asdasd")
	mes := message.SubMessage{
		Heard: message.Heard{
			MessageId:  "hefgbdferdvasdage",
			ProducerId: "asfqdacxasdqwcqwd",
			QueueId:    "c196de8f-d2d6-60d9-6092-c20433d156b3",
			Timestamp:  1687276300870,
			SendTime:   1687276300870,
			Indate:     0,
		},
		Subscriber: []string{"45a64421-35cc-6465-2046-3df88078542b"},
		Body:       []byte{2, 122, 98, 254},
	}
	message.SetCachedSubMessage(&mes)
	select {}
	//subMessage := message.GetCachedSubMessage("hefgbdfage")
	//fmt.Println("查询到的数据为", subMessage)
}

func TestType(t *testing.T) {
	mes := new(message.DelayedMessage)
	fmt.Println(getType(*mes))
}

func getType(mes any) int8 {
	switch mes.(type) {
	case message.RandomMessage:
		return enum.RandomMessage
	case message.DelayedMessage:
		return enum.DelayedMessage
	}
	return -1
}
func TestUUID(t *testing.T) {
	comm.Play.Add(1)
	comm.Play.Wait()
	//snowflake, err := utils.NewSnowflake(1)
	//if err != nil {
	//	return
	//}
	for i := 0; i < 10; i++ {
		id := utils.GetSnowflakeIdStr()
		fmt.Println(id)
		time.Sleep(100)
	}
}
