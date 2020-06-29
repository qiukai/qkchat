package filter

import (
	"encoding/json"
	"frame/bean"
	"net/http"
)

func _param(request *http.Request, controller bean.Controller) interface{} {
	length := request.ContentLength
	bytes := make([]byte, length)
	body := request.Body
	var param interface{}
	if body != nil {
		body.Read(bytes)
		param = controller.GetParam()
		err := json.Unmarshal(bytes, param)
		if err != nil {
			panic(err)
		}
	}
	return param
}
