package filter

import (
	"frame/config"
	"frame/filter"
	"my_error"
	"net/http"
)

type BlackNameFilter struct {
}

func (self *BlackNameFilter) DoFilter(response http.ResponseWriter, request *http.Request) error {

	value := config.GetConfig("uri_black_name")
	uri := request.RequestURI
	if value == uri {
		return my_error.NewError500ByMsg("黑名单uri")
	}
	return nil
}

func init() {
	filter.Register(&BlackNameFilter{})
}
