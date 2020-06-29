package filter

import (
	"frame/config"
	"frame/filter"
	"net/http"
)

type TraceIdFilter struct {
}

func (self *TraceIdFilter) DoFilter(response http.ResponseWriter, request *http.Request) error {
	config.GetTrace().AddLog()
	return nil
}

func init()  {
	filter.Register(&TraceIdFilter{})
}
