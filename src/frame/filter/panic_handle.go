package filter

import (
	"frame/myLog"
	"my_error"
	"net/http"
	"util"
)

func panicHanle(response http.ResponseWriter) {
	if err := recover(); err != nil {
		myLog.Info(err)
		_err,ok := err.(error)
		if ok {
			error500 := my_error.NewError500(_err)
			marshal := util.Marshal(error500)
			myLog.Error(string(marshal))
			result500(response, marshal)
			return
		}
		error500 := my_error.NewError500ByMsg("系统无法处理")
		marshal := util.Marshal(error500)
		result500(response, marshal)
	}
}
