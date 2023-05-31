package comm

import "waveQServer/comm/enum"

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
