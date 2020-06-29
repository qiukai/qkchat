package myLog

import (
	"frame/config"
	"log"
	"runtime"
	"runtime/debug"
)

var levelInt int

func Debug(v interface{}) {
	if 0 == levelInt {
		print(v, "debug")
	}
}

func Info(v interface{}) {
	if 1 < levelInt {
		return
	}
	print(v, "info")
}

func Warn(v interface{}) {
	if 2 < levelInt {
		return
	}
	print(v, "warn")
}

func Error(v interface{}) {
	if 3 < levelInt {
		return
	}
	print(v, "error")
}

func init() {
	level := config.GetConfig("log")
	Info("日志级别是：" + level)

	if "debug" == level {
		levelInt = 0
	} else if "info" == level {
		levelInt = 1
	} else if "warn" == level {
		levelInt = 2
	} else if "error" == level {
		levelInt = 3
	}

}

func print(v interface{}, level string) {
	traceId := config.GetTrace().GetTraceId()
	pc, _, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	if "error" == level {
		log.Print(level, " [", traceId, "] ", f.Name(), ":", line, " ", v, "\n", string(debug.Stack()))
	} else {
		log.Print(level, " [", traceId, "] ", f.Name(), ":", line, " ", v)
	}

}
