// v0.1.1
package mlg

import (
	"fmt"
	"time"
)

type MLManager struct {
	host string
	port string
}

func NewMLManager(port string) *MLManager {
	return &MLManager{
		host: "127.0.0.1",
		port: port,
	}
}

func (self *MLManager) Check () bool {
	resp := pureGet([]string{self.port, "check"})
	if resp=="okok"{
		return true
	} else {
		return false
	}
}

func (self *MLManager) Push (channel string, d string) string {
	dataMap := make(map[string]string)
	dataMap["data"] = d
	log := loopPost([]string{self.port, "push", channel}, dataMap, NETERRorTO)
	return log
}

func (self *MLManager) Get (channel string) string {
	resp := loopGet([]string{self.port, "get", channel}, NETERRorTO)
	return resp[4:]
}

func (self *MLManager) WaitForServer () {
	for !self.Check() {
		fmt.Println("服务器未启动")
	}
	time.Sleep(100*time.Millisecond)
}

func (self *MLManager) WaitForSignal(sig string, value string) {
	passReturn := "okok" + string(value)
	get := loopGet([]string{self.port, "getSignal", sig}, ONLY4NETERR)
	for get!= passReturn{
		time.Sleep(50*time.Millisecond)
		get = loopGet([]string{self.port, "getSignal", sig}, ONLY4NETERR)
		fmt.Println("得到一个不行的")
	}
}

func (self *MLManager) SetSignal (sig string, value string) string {
	dataMap := make(map[string]string)
	dataMap["data"] = value
	log := loopPost([]string{self.port, "setSignal", sig}, dataMap, ONLY4NETERR)
	return log
}









