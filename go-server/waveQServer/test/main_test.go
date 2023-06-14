package test

import (
	"fmt"
	"testing"
	"waveQServer/src/utils"
)

func TestMd5(t *testing.T) {
	md5 := utils.Md5([]byte("Admin"))
	fmt.Println(md5)
	t.Log()
}
