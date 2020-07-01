package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"my_error"
	"os"
	"runtime"
)

var configMap map[string]string

func GetConfig(key string) string {
	s := configMap[key]
	if "" == s {
		error500 := my_error.NewError500ByMsg("config.json 中没有 key:" + key)
		panic(error500)
	}
	return s
}

func init() {
	GetTrace().AddLog()
	dir, _ := os.Getwd()
	data, err := ioutil.ReadFile(dir + "/config/config.json")
	if err != nil {
		error500 := my_error.NewError500ByMsg("/config/config.json 文件不存在")
		print(error500)
		return
	}
	configMap = make(map[string]string)
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		panic(err)
	}

}

func print(v interface{}) {
	traceId := GetTrace().GetTraceId()
	pc, _, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	log.Print("info", " [", traceId, "] ", f.Name(), ":", line, " ", v)
}
