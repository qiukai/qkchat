package id

import (
	"time"
)

var ipStr string
var base = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var (
	currentTime    chan int64
	count          int64
	currentTimeNow int64
)

// 支持多机器 唯一id 生成器
func Next() string {
	num := make(chan string)
	go next(currentTime, num, false)
	return <-num
}

// 支持多机器 唯一id 生成器
func NextFormat() string {
	num := make(chan string)
	go next(currentTime, num, true)
	return <-num
}

func next(currentTime chan int64, num chan string, format bool) {
	now := <-currentTime
	if now == currentTimeNow {
		count++
	} else {
		count = 0
	}
	currentTimeNow = now
	now -= 1577808000

	var numTemp string
	if format {
		numTemp = ipStr + "-" + Int64ToString(now) + "-" + Int64ToString(count)
	} else {
		numTemp = ipStr + Int64ToString(now) + Int64ToString(count)
	}
	go lock()
	num <- numTemp
}

func lock() {
	currentTime <- getCuurTime()
}

func Int64ToString(value int64) string {
	var s string
	for q := value; q > 0; q = q / 62 {
		m := q % 62
		s = base[m] + s
	}

	if s == "" {
		s = "0"
	}

	return s
}

func getCuurTime() int64 {
	return time.Now().Unix()
}

func init() {
	count = 0
	currentTime = make(chan int64)
	currentTimeNow = getCuurTime()
	ipStr = Int64ToString(IpIntByDigit(2))
	go lock()
}
