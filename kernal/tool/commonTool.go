package tool

import (
	"github.com/axgle/mahonia"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Decode2str(b []byte, dec mahonia.Decoder)string{
	_, temp, _ := dec.Translate(b, true)
	return string(temp)
}

// ClearXmlinRun 将一个xml内容，清洗成一行一行的，这里可能会去掉一行最开始和最后的空格; 空号交给run cmd去解决吧
func ClearXmlinRun(s string) []string {
	ss := strings.Split(s, "\n")
	for i:=0;i<len(ss);i++ {
		reg := regexp.MustCompile("^[ \u3000\t\r\n]*")
		ss[i] = reg.ReplaceAllString(ss[i], "")

		reg = regexp.MustCompile("[ \u3000\t\r\n]*$")
		ss[i] = reg.ReplaceAllString(ss[i], "")
	}
	return ss
}


var ML_timeStartMs int64 = 1639238640520
// 获得int时间戳
func GetDiffTimeStampStr() string {
	i := time.Now().UnixMicro() - ML_timeStartMs
	return strconv.Itoa(int(i))
}

func BytesEquals(a,b []byte) bool {
	ret := true
	if len(a)!= len(b){
		ret = false
		return ret
	}
	for i:=0;i<len(a);i++{
		if a[i] != b[i]{
			ret = false
		}
	}
	return ret
}