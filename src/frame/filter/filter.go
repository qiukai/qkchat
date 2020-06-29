package filter

import (
	"net/http"
)


var filters []Filter

type Filter interface {
	DoFilter(response http.ResponseWriter, request *http.Request) error
}


// 执行所有的过滤器
func filter(response http.ResponseWriter, request *http.Request) bool {
	for _, value := range filters {
		err := value.DoFilter(response, request)
		if err != nil {
			response.Write([]byte(err.Error()))
			return false
		}
	}
	return true
}

func Register(filter Filter) {
	filters = append(filters, filter)
}