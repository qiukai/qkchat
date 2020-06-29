package myLog

import (
	"frame/config"
	"log"
)

var levelInt int

func Debug(v interface{}) {
	if 0 == levelInt {
		print(v)
	}
}

func Info(v interface{}) {
	if 1 < levelInt {
		return
	}
	print(v)
}

func Warn(v interface{}) {
	if 2 < levelInt {
		return
	}
	print(v)
}

func Error(v interface{}) {
	if 3 < levelInt {
		return
	}
	print(v)
}

func init() {
	level := config.GetConfig("log")
	print("日志级别是：" + level)

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

func print(v interface{}) {
	traceId := config.GetTrace().GetTraceId()
	log.Print("[", traceId, "] ", v)
}
