package errno

import "testing"

func TestBizError_Error(t *testing.T) {
	eo := NewSimpleBizError(ErrParameterInvalid, nil, "user_id")
	t.Log(eo.Error())

	eo = NewSimpleBizError(ErrMissingParameter, nil, "Code")
	t.Log(eo.Error())
}
