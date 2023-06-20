package test

import (
	"fmt"
	"reflect"
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
	var mes = make([]message.Message, 0, 5)
	subMessage := message.SubMessage{}
	randmoMes := message.RandomMessage{}
	mes = append(mes, &subMessage)
	mes = append(mes, &randmoMes)
	for _, v := range mes {
		//reflect.TypeOf(v).Kind()
		//randomMessage, ok := v.(*message.RandomMessage)
		//if ok {
		//
		//}
		fmt.Println("类型", reflect.TypeOf(v).Elem())
	}
}
