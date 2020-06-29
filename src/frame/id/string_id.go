package id

import (
	"time"
)

var ipInt int64
var base = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var (
	currentTime    chan int64
	count          int64
	currentTimeNow int64
)

func Next() string {
	num := make(chan string)
	go next(currentTime, num)
	return <-num
}

func next(currentTime chan int64, num chan string) {
	now := <-currentTime
	if now == currentTimeNow {
		count++
	} else {
		count = 1
	}
	currentTimeNow = now
	numTemp := int64ToString(ipInt) + int64ToString(now) + int64ToString(count)
	go lock()
	num <- numTemp
}

func lock() {
	currentTime <- getCuurTime()
}

func int64ToString(value int64) string {
	var s string
	for q := value; q > 0; q = q / 62 {
		m := q % 62
		s = base[m] + s
	}
	return s
}

func getCuurTime() int64 {
	return time.Now().Unix()
}

func init() {
	count = 1
	currentTime = make(chan int64)
	currentTimeNow = getCuurTime()
	ipInt = IpInt()
	go lock()
}


