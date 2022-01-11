package tool

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// GetPort todo 该函数缺乏测试，可能面临着端口已占用但是没提示的情况，请注意
func GetPort(port int) int {
	fmt.Println("当前测试的是端口", port)
	time.Sleep(1*time.Second)
	tl,err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err == nil {
		tl.Close()
		return port
	} else {
		return GetPort(port+1)
	}
}