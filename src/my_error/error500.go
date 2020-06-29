package my_error

type Error500 struct {
	ErrCode int
	ErrMsg  string
}

func NewError500(err error) *Error500 {
	return &Error500{ErrCode: 500, ErrMsg: err.Error()}
}

func NewError500ByMsg(ErrMsg string) *Error500 {
	return &Error500{ErrCode: 500, ErrMsg: ErrMsg}
}

func NewDefaultError500() *Error500 {
	return &Error500{ErrCode: 500, ErrMsg: "系统错误"}
}

func (self *Error500) Error() string {
	return self.ErrMsg
}
