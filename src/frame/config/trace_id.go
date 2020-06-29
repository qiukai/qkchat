package config

import (
	"frame/id"
	"util"
)

var logMap = make(map[int]string)

type Trace struct {

}

func GetTrace() *Trace {
	return new(Trace)
}

func (self *Trace) AddLog()  {
	goid := util.Goid()
	next := id.Next()
	logMap[goid] = next
}

func (self *Trace) GetTraceId() string {
	goid := util.Goid()
	return logMap[goid]
}
