package filter

import (
	"encoding/json"
	"net/http"
)

func _param(request *http.Request, param interface{}) interface{} {
	length := request.ContentLength
	bytes := make([]byte, length)
	body := request.Body
	if body != nil {
		body.Read(bytes)
		err := json.Unmarshal(bytes, param)
		if err != nil {
			panic(err)
		}
	}
	return param
}
