package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var configMap map[string]string

func GetConfig(key string) string {
	return configMap[key]
}

func init() {
	GetTrace().AddLog()
	dir, _ := os.Getwd()
	data, err := ioutil.ReadFile(dir + "/config/config.json")
	if err != nil {
		fmt.Println("/config/config.json 文件不存在", err)
		return
	}
	configMap = make(map[string]string)
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		panic(err)
	}

}
