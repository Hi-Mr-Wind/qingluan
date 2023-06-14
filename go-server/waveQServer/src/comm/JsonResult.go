package comm

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waveQServer/src/comm/enum"
)

type JsonResult struct {
	Code int `json:"code"`

	Mes string `json:"mes"`

	Data any
}

func NewJsonResult() *JsonResult {
	return new(JsonResult)
}

func OK(data ...interface{}) *JsonResult {
	result := NewJsonResult()
	result.Code = enum.OK
	result.Mes = "OK"
	result.Data = data
	return result
}
func Fail(mes string) *JsonResult {
	result := NewJsonResult()
	result.Code = enum.FAIL
	result.Mes = mes
	return result
}

// DisposeError 处理异常
func DisposeError(err error, c *gin.Context) {
	fail := Fail(err.Error())
	c.JSON(http.StatusBadRequest, fail)
	c.Abort()
	return
}
