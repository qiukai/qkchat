package my_error

import "util"

type Error404 struct {
	ErrCode int
	ErrMsg  string
}

func NewError404(msg string) *Error404 {
	return &Error404{ErrCode: 404, ErrMsg: msg}
}

func NewDefaultError404() *Error404 {
	return &Error404{ErrCode: 404, ErrMsg: "接口不存在"}
}

func (self *Error404) Error() string {
	return self.ErrMsg
}

func (self *Error404) String() string {
	return string(util.Marshal(self))
}
