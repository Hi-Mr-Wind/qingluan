package test

import (
	"fmt"
	"testing"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
	"waveQServer/src/utils"
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
