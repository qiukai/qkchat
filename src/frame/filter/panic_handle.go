package filter

import (
	"frame/myLog"
	"my_error"
	"net/http"
)

func panicHanle(response http.ResponseWriter) {
	if err := recover(); err != nil {
		_err, ok := err.(error)
		if ok {
			error500 := my_error.NewError500(_err)
			myLog.Error(error500)
			result500(response, []byte(error500.String()))
			return
		}

		s, ok := err.(string)
		if ok {
			error500 := my_error.NewError500ByMsg(s)
			myLog.Error(error500)
			result500(response, []byte(error500.String()))
			return
		}

		error500 := my_error.NewError500ByMsg("系统无法处理")
		myLog.Error(error500)
		result500(response, []byte(error500.String()))
	}
}
