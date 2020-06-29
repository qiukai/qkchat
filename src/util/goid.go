package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func Goid() int {
	defer func()  {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)        }
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	s := string(buf[:n])
	prefix := strings.TrimPrefix(s, "goroutine ")
	idField := strings.Fields(prefix)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

