package bean

import "net/http"

type Controller interface {
	DoService(response http.ResponseWriter, request *http.Request, value interface{}) interface{}

	GetParam() interface{}

}
