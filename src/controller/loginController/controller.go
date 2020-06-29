package loginController

import (
	"frame/bean"
	"frame/myLog"
	"net/http"
)

type LoginController struct {
}

func (self *LoginController) DoService(response http.ResponseWriter, request *http.Request, value interface{}) interface{} {
	param, _ := value.(*LoginParam)
	myLog.Info(param)
	return param
}

func (self *LoginController) GetParam() interface{} {
	return new(LoginParam)
}

func init() {
	bean.Register("/login", new(LoginController))
}
