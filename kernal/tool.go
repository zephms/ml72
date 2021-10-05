package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"
)

func showLogo(){
	fmt.Println("============================================\n")
	ascii := "              ____  _   __                 \n  __ _  __ __/ / /_(_) / /  ___ ____  ___ _\n /  ' \\/ // / / __/ / / /__/ _ `/ _ \\/ _ `/\n/_/_/_/\\_,_/_/\\__/_/ /____/\\_,_/_//_/\\_, / \n                                    /___/  \n"
	fmt.Println(ascii)
	fmt.Println("        multiLang   @  ml.acbs.top")
	fmt.Println("         Welcome to use MultiLang")
	fmt.Println("============================================\n")
}

func Checkproject(path string) bool {
	files,_ := ioutil.ReadDir(path)
	for _,i:= range files{
		if i.Name()=="project.xml"{
			return true
		}
	}
	return false
}

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

// IsExistDir IsDir todo 在当前权限访问不到的地方会有问题，比如访问 /root/abcd 会返回 true，但是实际上是不存在的
// IsDir todo 很遗憾的是，暂时没有找到好的区分新目录和随机字符串区分，所以目前的策略是禁止一切不是已有路径的访问。。。
func IsExistDir(pathstr string) bool {
	a, err := os.Stat(pathstr)
	if err!=nil{
		return false
	}
	return a.IsDir()
}

func Decode2str(b []byte, dec mahonia.Decoder)string{
	_, temp, _ := dec.Translate(b, true)
	return string(temp)
}