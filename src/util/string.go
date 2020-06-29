package util

import (
	"log"
	"strconv"
)


// 字符串转int64
func String2int64(str string) int64 {
	val1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val1
}

