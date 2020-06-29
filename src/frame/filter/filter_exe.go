package filter

import (
	_ "controller/loginController"
	"frame/bean"
	_ "frame/config"
	"my_error"
	"net/http"
	"util"
)



func Exe(response http.ResponseWriter, request *http.Request) {

	defer panicHanle(response)

	// 过滤器校验失败
	if !filter(response, request) {
		return
	}

	// 获取controller
	controller := bean.GetControllerByUri(request.RequestURI)

	if nil  == controller {
		panic(my_error.NewDefaultError404())
	}

	// 进行反序列化
	param := _param(request, controller)

	// 执行控制器
	service := controller.DoService(response, request, param)

	// 进行反序列化
	if service != nil {
		marshal := util.Marshal(service)
		result200(response, marshal)
		return
	}

	result200(response, []byte(""))
}




