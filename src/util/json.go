package util

import (
	"encoding/json"
	"reflect"
)

func Marshal(v interface{}) []byte {

	of := reflect.TypeOf(v)
	name := of.Name()
	b := name== "string"
	if b {
		return []byte(v.(string))
	}

	marshal, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return marshal
}
