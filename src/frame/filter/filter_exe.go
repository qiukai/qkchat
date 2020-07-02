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

	if nil == controller {
		panic(my_error.NewDefaultError404())
	}

	header := request.Header
	//[{"key":"Content-Type","value":"application/json","description":""}]
	contentTypeValue := header.Get("Content-Type")

	var param interface{}
	getParam := controller.GetParam()

	//typeOf := reflect.ValueOf(getParam)
	//i := typeOf.n
	//
	//myLog.Info(i)

	if "application/json" == contentTypeValue {
		// 进行反序列化
		param = _param(request, getParam)
	} else if "application/x-www-form-urlencoded;charset=UTF-8" == contentTypeValue {
		//form := request.Form
		//
		//
		//
		//form.
	}

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
