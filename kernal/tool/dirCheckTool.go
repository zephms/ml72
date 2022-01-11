package tool

import (
	"io/ioutil"
	"os"
)

func Checkproject(path string) bool {
	files,_ := ioutil.ReadDir(path)
	for _,i:= range files{
		if i.Name()=="project.xml"{
			return true
		}
	}
	return false
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