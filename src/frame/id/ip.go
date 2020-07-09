package id

import (
	"fmt"
	"net"
	"os"
	"strings"
	"util"
)

func Ip() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var ip string
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return ip
}

func IpInt() int64 {
	return IpIntByDigit(4)
}

func IpIntByDigit(digit uint8) int64 {
	ip := Ip()
	ipSplit := strings.Split(ip, ".")

	if digit == 1 {
		return util.String2int64(ipSplit[3])
	} else if digit == 2 {
		val3 := util.String2int64(ipSplit[2])
		val4 := util.String2int64(ipSplit[3])
		return val3*255 + val4
	} else if digit == 3 {
		val2 := util.String2int64(ipSplit[1])
		val3 := util.String2int64(ipSplit[2])
		val4 := util.String2int64(ipSplit[3])
		return val2*255*255 + val3*255 + val4
	} else if digit == 4 {
		val1 := util.String2int64(ipSplit[0])
		val2 := util.String2int64(ipSplit[1])
		val3 := util.String2int64(ipSplit[2])
		val4 := util.String2int64(ipSplit[3])
		sum := val1*255*255*255 + val2*255*255 + val3*255 + val4
		return sum
	} else {
		return 0
	}
}
